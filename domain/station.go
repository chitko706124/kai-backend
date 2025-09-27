package domain

import (
	"context"

	"github.com/LouisFernando1204/kai-backend.git/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Station struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name     string             `bson:"name" json:"name"`
	City     string             `bson:"city" json:"city"`
	Code     string             `bson:"code" json:"code"`
	IsActive bool               `bson:"is_active" json:"is_active"`
}

type StationRepository interface {
	FindAll(ctx context.Context) ([]Station, error)
	FindByID(ctx context.Context, id primitive.ObjectID) (*Station, error)
	Save(ctx context.Context, station *Station) error
	Update(ctx context.Context, station *Station) error
	Delete(ctx context.Context, id primitive.ObjectID) error
}

type StationService interface {
	GetAll(ctx context.Context) ([]dto.StationResponse, error)
	GetByID(ctx context.Context, id string) (*dto.StationResponse, error)
	Create(ctx context.Context, req dto.StationRequest) (*dto.StationResponse, error)
	Update(ctx context.Context, id string, req dto.StationRequest) (*dto.StationResponse, error)
	Delete(ctx context.Context, id string) error
}
