package repository

import (
	"context"
	"errors"

	"github.com/LouisFernando1204/kai-backend.git/domain"
	"github.com/LouisFernando1204/kai-backend.git/internal/config"
	"github.com/LouisFernando1204/kai-backend.git/internal/connection"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	Collection *mongo.Collection
}

func NewUserRepository(conf config.Database) domain.UserRepository {
	return &userRepository{
		Collection: connection.GetCollection(conf.MongoDatabase, "users"),
	}
}

func (r *userRepository) Save(ctx context.Context, user *domain.User) error {
	_, err := r.Collection.InsertOne(ctx, user)
	return err
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	err := r.Collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}
