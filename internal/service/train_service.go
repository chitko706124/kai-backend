package service

import (
	"context"
	"errors"

	"github.com/LouisFernando1204/kai-backend.git/domain"
	"github.com/LouisFernando1204/kai-backend.git/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type trainService struct {
	trainRepository domain.TrainRepository
}

func NewTrainService(trainRepo domain.TrainRepository) domain.TrainService {
	return &trainService{
		trainRepository: trainRepo,
	}
}

func carriagesDTOToDomain(dtoCarriages []dto.CarriageDTO) []domain.Carriage {
	var domainCarriages []domain.Carriage
	for _, c := range dtoCarriages {
		var domainSeats []domain.Seat
		for _, s := range c.Seats {
			domainSeats = append(domainSeats, domain.Seat{
				Code:        s.Code,
				IsAvailable: s.IsAvailable,
			})
		}
		domainCarriages = append(domainCarriages, domain.Carriage{
			Code:  c.Code,
			Seats: domainSeats,
		})
	}
	return domainCarriages
}

func carriagesDomainToDTO(domainCarriages []domain.Carriage) []dto.CarriageDTO {
	var dtoCarriages []dto.CarriageDTO
	for _, c := range domainCarriages {
		var dtoSeats []dto.SeatDTO
		for _, s := range c.Seats {
			dtoSeats = append(dtoSeats, dto.SeatDTO{
				Code:        s.Code,
				IsAvailable: s.IsAvailable,
			})
		}
		dtoCarriages = append(dtoCarriages, dto.CarriageDTO{
			Code:  c.Code,
			Seats: dtoSeats,
		})
	}
	return dtoCarriages
}

func (s *trainService) GetAll(ctx context.Context) ([]dto.TrainResponse, error) {
	trains, err := s.trainRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var trainResponses []dto.TrainResponse
	for _, train := range trains {
		trainResponses = append(trainResponses, dto.TrainResponse{
			ID:        train.ID.Hex(),
			Name:      train.Name,
			Class:     train.Class,
			Carriages: carriagesDomainToDTO(train.Carriages),
		})
	}
	return trainResponses, nil
}

func (s *trainService) GetByID(ctx context.Context, id string) (*dto.TrainResponse, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid id format")
	}

	train, err := s.trainRepository.FindByID(ctx, objectID)
	if err != nil {
		return nil, err
	}

	response := &dto.TrainResponse{
		ID:        train.ID.Hex(),
		Name:      train.Name,
		Class:     train.Class,
		Carriages: carriagesDomainToDTO(train.Carriages),
	}
	return response, nil
}

func (s *trainService) Create(ctx context.Context, req dto.CreateTrainRequest) (*dto.TrainResponse, error) {
	newTrain := &domain.Train{
		ID:        primitive.NewObjectID(),
		Name:      req.Name,
		Class:     req.Class,
		Carriages: carriagesDTOToDomain(req.Carriages),
	}

	if err := s.trainRepository.Save(ctx, newTrain); err != nil {
		return nil, err
	}

	response := &dto.TrainResponse{
		ID:        newTrain.ID.Hex(),
		Name:      newTrain.Name,
		Class:     newTrain.Class,
		Carriages: carriagesDomainToDTO(newTrain.Carriages),
	}
	return response, nil
}

func (s *trainService) Update(ctx context.Context, id string, req dto.UpdateTrainRequest) (*dto.TrainResponse, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid id format")
	}

	existingTrain, err := s.trainRepository.FindByID(ctx, objectID)
	if err != nil {
		return nil, err
	}

	existingTrain.Name = req.Name
	existingTrain.Class = req.Class
	existingTrain.Carriages = carriagesDTOToDomain(req.Carriages)

	if err := s.trainRepository.Update(ctx, existingTrain); err != nil {
		return nil, err
	}

	response := &dto.TrainResponse{
		ID:        existingTrain.ID.Hex(),
		Name:      existingTrain.Name,
		Class:     existingTrain.Class,
		Carriages: carriagesDomainToDTO(existingTrain.Carriages),
	}
	return response, nil
}

func (s *trainService) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid id format")
	}
	return s.trainRepository.Delete(ctx, objectID)
}
