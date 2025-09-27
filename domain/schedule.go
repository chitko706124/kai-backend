package domain

import (
	"context"
	"time"

	"github.com/LouisFernando1204/kai-backend.git/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Schedule struct {
	ID                   primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	TrainID              primitive.ObjectID `bson:"train_id" json:"train_id"`
	OriginStationID      primitive.ObjectID `bson:"origin_station_id" json:"origin_station_id"`
	DestinationStationID primitive.ObjectID `bson:"destination_station_id" json:"destination_station_id"`
	DepartureTime        time.Time          `bson:"departure_time" json:"departure_time"`
	ArrivalTime          time.Time          `bson:"arrival_time" json:"arrival_time"`
	Price                float64            `bson:"price" json:"price"`
	AvailableSeats       int                `bson:"available_seats" json:"available_seats"`
}

type ScheduleRepository interface {
	FindAll(ctx context.Context) ([]Schedule, error)
	FindByID(ctx context.Context, id primitive.ObjectID) (*Schedule, error)
	Save(ctx context.Context, schedule *Schedule) error
	Update(ctx context.Context, schedule *Schedule) error
	Delete(ctx context.Context, id primitive.ObjectID) error
	Search(ctx context.Context, originID, destID primitive.ObjectID, date time.Time) ([]Schedule, error)
}

type ScheduleService interface {
	GetAll(ctx context.Context) ([]dto.ScheduleResponse, error)
	GetByID(ctx context.Context, id string) (*dto.ScheduleResponse, error)
	GetSeatLayout(ctx context.Context, scheduleID string) ([]dto.CarriageLayoutDTO, error)
	Create(ctx context.Context, req dto.CreateScheduleRequest) (*dto.ScheduleResponse, error)
	Update(ctx context.Context, id string, req dto.UpdateScheduleRequest) (*dto.ScheduleResponse, error)
	Delete(ctx context.Context, id string) error
	Search(ctx context.Context, req dto.SearchScheduleRequest) ([]dto.ScheduleResponse, error)
}
