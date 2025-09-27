package service

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/LouisFernando1204/kai-backend.git/domain"
	"github.com/LouisFernando1204/kai-backend.git/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const ADMIN_FEE = 7500.00

type bookingService struct {
	bookingRepository  domain.BookingRepository
	scheduleRepository domain.ScheduleRepository
	scheduleService    domain.ScheduleService
}

func NewBookingService(
	bookingRepo domain.BookingRepository,
	scheduleRepo domain.ScheduleRepository,
	scheduleSvc domain.ScheduleService,
) domain.BookingService {
	return &bookingService{
		bookingRepository:  bookingRepo,
		scheduleRepository: scheduleRepo,
		scheduleService:    scheduleSvc,
	}
}

func generateBookingCode() string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var sb strings.Builder
	sb.Grow(8)
	for i := 0; i < 8; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	return sb.String()
}

func (s *bookingService) GetBookingByID(ctx context.Context, id string) (*dto.BookingResponse, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid id format")
	}

	booking, err := s.bookingRepository.FindByID(ctx, objID)
	if err != nil {
		return nil, err
	}

	scheduleResponse, err := s.scheduleService.GetByID(ctx, booking.ScheduleID.Hex())
	if err != nil {
		return nil, err
	}

	var passengers []dto.PassengerRequest
	for _, p := range booking.Passengers {
		passengers = append(passengers, dto.PassengerRequest{
			FullName:       p.FullName,
			IdentityNumber: p.IdentityNumber,
		})
	}

	return &dto.BookingResponse{
		ID:            booking.ID.Hex(),
		BookingCode:   booking.BookingCode,
		Schedule:      *scheduleResponse,
		Passengers:    passengers,
		TotalPrice:    booking.TotalPrice,
		Status:        booking.Status,
		PaymentExpiry: booking.PaymentExpiry.Format("2006-01-02 15:04"),
	}, nil
}

func (s *bookingService) GetBookingsByUserID(ctx context.Context, userID string) ([]dto.BookingResponse, error) {
	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, errors.New("invalid user id")
	}

	bookings, err := s.bookingRepository.FindByUserID(ctx, userObjID)
	if err != nil {
		return nil, err
	}

	var responses []dto.BookingResponse
	for _, booking := range bookings {
		response, err := s.GetBookingByID(ctx, booking.ID.Hex())
		if err != nil {
			fmt.Printf("Error building response for booking %s: %v\n", booking.ID.Hex(), err)
			continue
		}
		responses = append(responses, *response)
	}

	return responses, nil
}

func (s *bookingService) CreateBooking(ctx context.Context, userID string, req dto.CreateBookingRequest) (*dto.BookingResponse, error) {
	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, errors.New("invalid user id")
	}
	scheduleObjID, err := primitive.ObjectIDFromHex(req.ScheduleID)
	if err != nil {
		return nil, errors.New("invalid schedule id")
	}

	schedule, err := s.scheduleRepository.FindByID(ctx, scheduleObjID)
	if err != nil {
		return nil, err
	}
	if schedule.AvailableSeats < len(req.Passengers) {
		return nil, errors.New("insufficient available seats")
	}

	totalPrice := (schedule.Price * float64(len(req.Passengers))) + ADMIN_FEE

	var passengers []domain.Passenger
	for _, p := range req.Passengers {
		passengers = append(passengers, domain.Passenger{
			FullName:       p.FullName,
			IdentityNumber: p.IdentityNumber,
			Seat:           p.SeatCode,
		})
	}

	newBooking := &domain.Booking{
		ID:          primitive.NewObjectID(),
		UserID:      userObjID,
		ScheduleID:  scheduleObjID,
		BookingCode: generateBookingCode(),
		Booker: domain.BookerInfo{
			FullName:    req.Booker.FullName,
			Email:       req.Booker.Email,
			PhoneNumber: req.Booker.PhoneNumber,
		},
		Passengers:    passengers,
		TotalPrice:    totalPrice,
		AdminFee:      ADMIN_FEE,
		Status:        "PENDING_PAYMENT",
		PaymentExpiry: time.Now().Add(30 * time.Minute),
		CreatedAt:     time.Now(),
	}

	if err := s.bookingRepository.Save(ctx, newBooking); err != nil {
		return nil, err
	}

	schedule.AvailableSeats -= len(req.Passengers)
	if err := s.scheduleRepository.Update(ctx, schedule); err != nil {
		fmt.Printf("CRITICAL: Failed to update available seats for schedule %s after booking %s\n", schedule.ID.Hex(), newBooking.ID.Hex())
	}

	scheduleResponse, err := s.scheduleService.GetByID(ctx, schedule.ID.Hex())
	if err != nil {
		return nil, err
	}

	return &dto.BookingResponse{
		ID:            newBooking.ID.Hex(),
		BookingCode:   newBooking.BookingCode,
		Schedule:      *scheduleResponse,
		Passengers:    req.Passengers,
		TotalPrice:    newBooking.TotalPrice,
		Status:        newBooking.Status,
		PaymentExpiry: newBooking.PaymentExpiry.Format("2006-01-02 15:04"),
	}, nil
}

func (s *bookingService) UpdateBookingStatus(ctx context.Context, id string, req dto.UpdateBookingStatusRequest) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid id format")
	}

	booking, err := s.bookingRepository.FindByID(ctx, objID)
	if err != nil {
		return err
	}

	if booking.Status == "PAID" || booking.Status == "CANCELLED" {
		return fmt.Errorf("booking status cannot be changed from %s", booking.Status)
	}

	return s.bookingRepository.UpdateStatus(ctx, objID, req.Status)
}

func (s *bookingService) DeleteBooking(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid id format")
	}
	return s.bookingRepository.Delete(ctx, objID)
}

func (s *bookingService) CancelBooking(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid id format")
	}

	booking, err := s.bookingRepository.FindByID(ctx, objID)
	if err != nil {
		return err
	}

	schedule, err := s.scheduleRepository.FindByID(ctx, booking.ScheduleID)
	if err != nil {
		return err
	}

	schedule.AvailableSeats += len(booking.Passengers)
	if err := s.scheduleRepository.Update(ctx, schedule); err != nil {
		return err
	}

	return s.bookingRepository.UpdateStatus(ctx, objID, "CANCELLED")
}
