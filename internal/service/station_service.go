package service

import (
	"context"
	"errors"

	"github.com/LouisFernando1204/kai-backend.git/domain"
	"github.com/LouisFernando1204/kai-backend.git/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type stationService struct {
	stationRepository domain.StationRepository
}

func NewStationService(stationRepo domain.StationRepository) domain.StationService {
	return &stationService{
		stationRepository: stationRepo,
	}
}

func (s *stationService) GetAll(ctx context.Context) ([]dto.StationResponse, error) {
	stations, err := s.stationRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var stationResponses []dto.StationResponse
	for _, station := range stations {
		stationResponses = append(stationResponses, dto.StationResponse{
			ID:   station.ID.Hex(),
			Name: station.Name,
			City: station.City,
			Code: station.Code,
		})
	}

	return stationResponses, nil
}

func (s *stationService) GetByID(ctx context.Context, id string) (*dto.StationResponse, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid id format")
	}

	station, err := s.stationRepository.FindByID(ctx, objectID)
	if err != nil {
		return nil, err
	}

	response := &dto.StationResponse{
		ID:   station.ID.Hex(),
		Name: station.Name,
		City: station.City,
		Code: station.Code,
	}

	return response, nil
}

func (s *stationService) Create(ctx context.Context, req dto.StationRequest) (*dto.StationResponse, error) {
	newStation := &domain.Station{
		ID:       primitive.NewObjectID(),
		Name:     req.Name,
		City:     req.City,
		Code:     req.Code,
		IsActive: true,
	}

	err := s.stationRepository.Save(ctx, newStation)
	if err != nil {
		return nil, err
	}

	response := &dto.StationResponse{
		ID:   newStation.ID.Hex(),
		Name: newStation.Name,
		City: newStation.City,
		Code: newStation.Code,
	}

	return response, nil
}

func (s *stationService) Update(ctx context.Context, id string, req dto.StationRequest) (*dto.StationResponse, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid id format")
	}

	existingStation, err := s.stationRepository.FindByID(ctx, objectID)
	if err != nil {
		return nil, err
	}

	existingStation.Name = req.Name
	existingStation.City = req.City
	existingStation.Code = req.Code

	err = s.stationRepository.Update(ctx, existingStation)
	if err != nil {
		return nil, err
	}

	response := &dto.StationResponse{
		ID:   existingStation.ID.Hex(),
		Name: existingStation.Name,
		City: existingStation.City,
		Code: existingStation.Code,
	}

	return response, nil
}

func (s *stationService) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid id format")
	}

	return s.stationRepository.Delete(ctx, objectID)
}
