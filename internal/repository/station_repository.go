package repository

import (
	"context"
	"errors"

	"github.com/LouisFernando1204/kai-backend.git/domain"
	"github.com/LouisFernando1204/kai-backend.git/internal/config"
	"github.com/LouisFernando1204/kai-backend.git/internal/connection"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type stationRepository struct {
	Collection *mongo.Collection
}

func NewStationRepository(conf config.Database) domain.StationRepository {
	return &stationRepository{
		Collection: connection.GetCollection(conf.MongoDatabase, "stations"),
	}
}

func (r *stationRepository) FindAll(ctx context.Context) ([]domain.Station, error) {
	var stations []domain.Station

	cursor, err := r.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &stations); err != nil {
		return nil, err
	}

	return stations, nil
}

func (r *stationRepository) FindByID(ctx context.Context, id primitive.ObjectID) (*domain.Station, error) {
	var station domain.Station
	err := r.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&station)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("station not found")
		}
		return nil, err
	}

	return &station, nil
}

func (r *stationRepository) Save(ctx context.Context, station *domain.Station) error {
	_, err := r.Collection.InsertOne(ctx, station)
	return err
}

func (r *stationRepository) Update(ctx context.Context, station *domain.Station) error {
	filter := bson.M{"_id": station.ID}
	update := bson.M{"$set": station}

	_, err := r.Collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *stationRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	_, err := r.Collection.DeleteOne(ctx, filter)
	return err
}
