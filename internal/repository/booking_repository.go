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

type bookingRepository struct {
	Collection *mongo.Collection
}

func NewBookingRepository(conf config.Database) domain.BookingRepository {
	return &bookingRepository{
		Collection: connection.GetCollection(conf.MongoDatabase, "bookings"),
	}
}

func (r *bookingRepository) FindAll(ctx context.Context) ([]domain.Booking, error) {
	var bookings []domain.Booking

	cursor, err := r.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &bookings); err != nil {
		return nil, err
	}

	return bookings, nil
}

func (r *bookingRepository) FindByID(ctx context.Context, id primitive.ObjectID) (*domain.Booking, error) {
	var booking domain.Booking
	err := r.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&booking)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("booking not found")
		}
		return nil, err
	}
	return &booking, nil
}

func (r *bookingRepository) FindByUserID(ctx context.Context, userID primitive.ObjectID) ([]domain.Booking, error) {
	var bookings []domain.Booking

	cursor, err := r.Collection.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &bookings); err != nil {
		return nil, err
	}

	return bookings, nil
}

func (r *bookingRepository) FindByScheduleID(ctx context.Context, scheduleID primitive.ObjectID) ([]domain.Booking, error) {
	var bookings []domain.Booking
	cursor, err := r.Collection.Find(ctx, bson.M{"schedule_id": scheduleID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	if err = cursor.All(ctx, &bookings); err != nil {
		return nil, err
	}
	return bookings, nil
}

func (r *bookingRepository) Save(ctx context.Context, booking *domain.Booking) error {
	_, err := r.Collection.InsertOne(ctx, booking)
	return err
}

func (r *bookingRepository) Update(ctx context.Context, booking *domain.Booking) error {
	filter := bson.M{"_id": booking.ID}
	update := bson.M{"$set": booking}

	_, err := r.Collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *bookingRepository) UpdateStatus(ctx context.Context, id primitive.ObjectID, status string) error {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"status": status,
		},
	}

	result, err := r.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.ModifiedCount == 0 {
		return errors.New("no booking was updated, check if the ID is correct")
	}

	return nil
}

func (r *bookingRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	_, err := r.Collection.DeleteOne(ctx, filter)
	return err
}
