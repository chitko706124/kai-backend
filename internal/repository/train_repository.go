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

type trainRepository struct {
	Collection *mongo.Collection
}

func NewTrainRepository(conf config.Database) domain.TrainRepository {
	return &trainRepository{
		Collection: connection.GetCollection(conf.MongoDatabase, "trains"),
	}
}

func (r *trainRepository) FindAll(ctx context.Context) ([]domain.Train, error) {
	var trains []domain.Train

	cursor, err := r.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &trains); err != nil {
		return nil, err
	}

	return trains, nil
}

func (r *trainRepository) FindByID(ctx context.Context, id primitive.ObjectID) (*domain.Train, error) {
	var train domain.Train
	err := r.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&train)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("train not found")
		}
		return nil, err
	}

	return &train, nil
}

func (r *trainRepository) Save(ctx context.Context, train *domain.Train) error {
	_, err := r.Collection.InsertOne(ctx, train)
	return err
}

func (r *trainRepository) Update(ctx context.Context, train *domain.Train) error {
	filter := bson.M{"_id": train.ID}
	update := bson.M{"$set": train}
	_, err := r.Collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *trainRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	filter := bson.M{"_id": id}
	_, err := r.Collection.DeleteOne(ctx, filter)
	return err
}
