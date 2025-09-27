package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FullName       string             `bson:"full_name" json:"full_name"`
	Email          string             `bson:"email" json:"email"`
	Password       string             `bson:"password" json:"-"`
	PhoneNumber    string             `bson:"phone_number" json:"phone_number"`
	IdentityNumber string             `bson:"identity_number" json:"identity_number"` // No KTP
	CreatedAt      time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt      time.Time          `bson:"updated_at" json:"updated_at"`
}

type UserRepository interface {
	Save(ctx context.Context, user *User) error
	FindByEmail(ctx context.Context, email string) (*User, error)
}
