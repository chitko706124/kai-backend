package domain

import (
	"context"

	"github.com/LouisFernando1204/kai-backend.git/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Seat struct {
	Code        string `bson:"code" json:"code"`
	IsAvailable bool   `bson:"is_available" json:"is_available"`
}

type Carriage struct {
	Code  string `bson:"code" json:"code"`
	Seats []Seat `bson:"seats" json:"seats"`
}

type Train struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Class     string             `bson:"class" json:"class"`
	Carriages []Carriage         `bson:"carriages" json:"carriages"`
}

type TrainRepository interface {
	FindAll(ctx context.Context) ([]Train, error)
	FindByID(ctx context.Context, id primitive.ObjectID) (*Train, error)
	Save(ctx context.Context, train *Train) error
	Update(ctx context.Context, train *Train) error
	Delete(ctx context.Context, id primitive.ObjectID) error
}

type TrainService interface {
	GetAll(ctx context.Context) ([]dto.TrainResponse, error)
	GetByID(ctx context.Context, id string) (*dto.TrainResponse, error)
	Create(ctx context.Context, req dto.CreateTrainRequest) (*dto.TrainResponse, error)
	Update(ctx context.Context, id string, req dto.UpdateTrainRequest) (*dto.TrainResponse, error)
	Delete(ctx context.Context, id string) error
}
