package repository

import (
	"context"
	"errors"
	"time"

	"github.com/LouisFernando1204/kai-backend.git/domain"
	"github.com/LouisFernando1204/kai-backend.git/internal/config"
	"github.com/LouisFernando1204/kai-backend.git/internal/connection"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type scheduleRepository struct {
	Collection *mongo.Collection
}

func NewScheduleRepository(conf config.Database) domain.ScheduleRepository {
	return &scheduleRepository{
		Collection: connection.GetCollection(conf.MongoDatabase, "schedules"),
	}
}

func (r *scheduleRepository) FindAll(ctx context.Context) ([]domain.Schedule, error) {
	var schedules []domain.Schedule

	cursor, err := r.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &schedules); err != nil {
		return nil, err
	}

	return schedules, nil
}

func (r *scheduleRepository) FindByID(ctx context.Context, id primitive.ObjectID) (*domain.Schedule, error) {
	var schedule domain.Schedule
	err := r.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&schedule)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("schedule not found")
		}
		return nil, err
	}
	return &schedule, nil
}

func (r *scheduleRepository) Save(ctx context.Context, schedule *domain.Schedule) error {
	_, err := r.Collection.InsertOne(ctx, schedule)
	return err
}

func (r *scheduleRepository) Update(ctx context.Context, schedule *domain.Schedule) error {
	filter := bson.M{"_id": schedule.ID}
	update := bson.M{"$set": schedule}
	_, err := r.Collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *scheduleRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	_, err := r.Collection.DeleteOne(ctx, filter)
	return err
}

func (r *scheduleRepository) Search(ctx context.Context, originID primitive.ObjectID, destID primitive.ObjectID, date time.Time) ([]domain.Schedule, error) {
	var schedules []domain.Schedule

	startOfDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	endOfDay := startOfDay.Add(24 * time.Hour)

	filter := bson.M{
		"origin_station_id":      originID,
		"destination_station_id": destID,
		"departure_time": bson.M{
			"$gte": startOfDay,
			"$lt":  endOfDay,
		},
	}

	cursor, err := r.Collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &schedules); err != nil {
		return nil, err
	}

	return schedules, nil
}
