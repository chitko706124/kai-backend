package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/LouisFernando1204/kai-backend.git/domain"
	"github.com/LouisFernando1204/kai-backend.git/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type scheduleService struct {
	scheduleRepository domain.ScheduleRepository
	trainRepository    domain.TrainRepository
	stationRepository  domain.StationRepository
	bookingRepository  domain.BookingRepository
}

func NewScheduleService(
	scheduleRepo domain.ScheduleRepository,
	trainRepo domain.TrainRepository,
	stationRepo domain.StationRepository,
	bookingRepo domain.BookingRepository,
) domain.ScheduleService {
	return &scheduleService{
		scheduleRepository: scheduleRepo,
		trainRepository:    trainRepo,
		stationRepository:  stationRepo,
		bookingRepository:  bookingRepo,
	}
}

func (s *scheduleService) buildScheduleResponse(ctx context.Context, schedule domain.Schedule) (*dto.ScheduleResponse, error) {
	origin, err := s.stationRepository.FindByID(ctx, schedule.OriginStationID)
	if err != nil {
		return nil, err
	}
	destination, err := s.stationRepository.FindByID(ctx, schedule.DestinationStationID)
	if err != nil {
		return nil, err
	}
	train, err := s.trainRepository.FindByID(ctx, schedule.TrainID)
	if err != nil {
		return nil, err
	}

	duration := schedule.ArrivalTime.Sub(schedule.DepartureTime)
	durationStr := fmt.Sprintf("%dh %dm", int(duration.Hours()), int(duration.Minutes())%60)

	return &dto.ScheduleResponse{
		ID: schedule.ID.Hex(),
		Train: dto.TrainInScheduleResponse{
			Name:  train.Name,
			Class: train.Class,
		},
		OriginStation: dto.StationResponse{
			ID:   origin.ID.Hex(),
			Name: origin.Name,
			City: origin.City,
			Code: origin.Code,
		},
		DestinationStation: dto.StationResponse{
			ID:   destination.ID.Hex(),
			Name: destination.Name,
			City: destination.City,
			Code: destination.Code,
		},
		DepartureTime:  schedule.DepartureTime.Format("2006-01-02 15:04"),
		ArrivalTime:    schedule.ArrivalTime.Format("2006-01-02 15:04"),
		Duration:       durationStr,
		Price:          schedule.Price,
		AvailableSeats: schedule.AvailableSeats,
	}, nil
}

func (s *scheduleService) GetAll(ctx context.Context) ([]dto.ScheduleResponse, error) {
	schedules, err := s.scheduleRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var responses []dto.ScheduleResponse
	for _, schedule := range schedules {
		response, err := s.buildScheduleResponse(ctx, schedule)
		if err != nil {
			fmt.Printf("Error building response for schedule %s: %v\n", schedule.ID.Hex(), err)
			continue
		}
		responses = append(responses, *response)
	}

	return responses, nil
}

func (s *scheduleService) GetByID(ctx context.Context, id string) (*dto.ScheduleResponse, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid id format")
	}

	schedule, err := s.scheduleRepository.FindByID(ctx, objectID)
	if err != nil {
		return nil, err
	}

	return s.buildScheduleResponse(ctx, *schedule)
}

func (s *scheduleService) GetSeatLayout(ctx context.Context, scheduleID string) ([]dto.CarriageLayoutDTO, error) {
	scheduleObjID, err := primitive.ObjectIDFromHex(scheduleID)
	if err != nil {
		return nil, errors.New("invalid schedule id")
	}

	schedule, err := s.scheduleRepository.FindByID(ctx, scheduleObjID)
	if err != nil {
		return nil, err
	}

	train, err := s.trainRepository.FindByID(ctx, schedule.TrainID)
	if err != nil {
		return nil, err
	}

	bookings, err := s.bookingRepository.FindByScheduleID(ctx, scheduleObjID)
	if err != nil {
		return nil, err
	}

	takenSeats := make(map[string]bool)
	for _, booking := range bookings {
		for _, passenger := range booking.Passengers {
			if passenger.Seat != "" {
				takenSeats[passenger.Seat] = true
			}
		}
	}

	var layout []dto.CarriageLayoutDTO
	for _, carriage := range train.Carriages {
		var seatLayouts []dto.SeatAvailabilityDTO
		for _, seat := range carriage.Seats {
			isAvailable := !takenSeats[seat.Code]
			seatLayouts = append(seatLayouts, dto.SeatAvailabilityDTO{
				Code:        seat.Code,
				IsAvailable: isAvailable,
			})
		}
		layout = append(layout, dto.CarriageLayoutDTO{
			Code:  carriage.Code,
			Seats: seatLayouts,
		})
	}

	return layout, nil
}

func (s *scheduleService) Create(ctx context.Context, req dto.CreateScheduleRequest) (*dto.ScheduleResponse, error) {
	trainID, _ := primitive.ObjectIDFromHex(req.TrainID)
	originID, _ := primitive.ObjectIDFromHex(req.OriginStationID)
	destID, _ := primitive.ObjectIDFromHex(req.DestinationStationID)
	departureTime, _ := time.Parse("2006-01-02 15:04", req.DepartureTime)
	arrivalTime, _ := time.Parse("2006-01-02 15:04", req.ArrivalTime)

	newSchedule := &domain.Schedule{
		ID:                   primitive.NewObjectID(),
		TrainID:              trainID,
		OriginStationID:      originID,
		DestinationStationID: destID,
		DepartureTime:        departureTime,
		ArrivalTime:          arrivalTime,
		Price:                req.Price,
		AvailableSeats:       req.AvailableSeats,
	}

	if err := s.scheduleRepository.Save(ctx, newSchedule); err != nil {
		return nil, err
	}

	return s.buildScheduleResponse(ctx, *newSchedule)
}

func (s *scheduleService) Update(ctx context.Context, id string, req dto.UpdateScheduleRequest) (*dto.ScheduleResponse, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid id format")
	}

	existingSchedule, err := s.scheduleRepository.FindByID(ctx, objectID)
	if err != nil {
		return nil, err
	}

	departureTime, _ := time.Parse("2006-01-02 15:04", req.DepartureTime)
	arrivalTime, _ := time.Parse("2006-01-02 15:04", req.ArrivalTime)

	existingSchedule.DepartureTime = departureTime
	existingSchedule.ArrivalTime = arrivalTime
	existingSchedule.Price = req.Price
	existingSchedule.AvailableSeats = req.AvailableSeats

	if err := s.scheduleRepository.Update(ctx, existingSchedule); err != nil {
		return nil, err
	}

	return s.buildScheduleResponse(ctx, *existingSchedule)
}

func (s *scheduleService) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid id format")
	}
	return s.scheduleRepository.Delete(ctx, objectID)
}

func (s *scheduleService) Search(ctx context.Context, req dto.SearchScheduleRequest) ([]dto.ScheduleResponse, error) {
	originID, err := primitive.ObjectIDFromHex(req.OriginStationID)
	if err != nil {
		return nil, errors.New("invalid origin station id")
	}
	destID, err := primitive.ObjectIDFromHex(req.DestinationStationID)
	if err != nil {
		return nil, errors.New("invalid destination station id")
	}
	date, err := time.Parse("2006-01-02", req.DepartureDate)
	if err != nil {
		return nil, errors.New("invalid date format, use YYYY-MM-DD")
	}

	schedules, err := s.scheduleRepository.Search(ctx, originID, destID, date)
	if err != nil {
		return nil, err
	}

	var responses []dto.ScheduleResponse
	for _, schedule := range schedules {
		if schedule.AvailableSeats >= req.Passengers {
			response, err := s.buildScheduleResponse(ctx, schedule)
			if err != nil {
				fmt.Printf("Error building response for schedule %s: %v\n", schedule.ID.Hex(), err)
				continue
			}
			responses = append(responses, *response)
		}
	}

	return responses, nil
}
