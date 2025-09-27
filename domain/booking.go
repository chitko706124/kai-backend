package domain

import (
	"context"
	"time"

	"github.com/LouisFernando1204/kai-backend.git/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookerInfo struct {
	FullName    string `bson:"full_name" json:"full_name"`
	Email       string `bson:"email" json:"email"`
	PhoneNumber string `bson:"phone_number" json:"phone_number"`
}

type Passenger struct {
	FullName       string `bson:"full_name" json:"full_name"`
	IdentityNumber string `bson:"identity_number" json:"identity_number"`
	Seat           string `bson:"seat" json:"seat"`
}

type Booking struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID        primitive.ObjectID `bson:"user_id" json:"user_id"`
	ScheduleID    primitive.ObjectID `bson:"schedule_id" json:"schedule_id"`
	BookingCode   string             `bson:"booking_code" json:"booking_code"`
	Booker        BookerInfo         `bson:"booker" json:"booker"`
	Passengers    []Passenger        `bson:"passengers" json:"passengers"`
	TotalPrice    float64            `bson:"total_price" json:"total_price"`
	AdminFee      float64            `bson:"admin_fee" json:"admin_fee"`
	Status        string             `bson:"status" json:"status"`
	PaymentExpiry time.Time          `bson:"payment_expiry" json:"payment_expiry"`
	CreatedAt     time.Time          `bson:"created_at" json:"created_at"`
}

type BookingRepository interface {
	FindAll(ctx context.Context) ([]Booking, error)
	FindByID(ctx context.Context, id primitive.ObjectID) (*Booking, error)
	FindByUserID(ctx context.Context, userID primitive.ObjectID) ([]Booking, error)
	FindByScheduleID(ctx context.Context, scheduleID primitive.ObjectID) ([]Booking, error)
	Save(ctx context.Context, booking *Booking) error
	Update(ctx context.Context, booking *Booking) error
	UpdateStatus(ctx context.Context, id primitive.ObjectID, status string) error
	Delete(ctx context.Context, id primitive.ObjectID) error
}

type BookingService interface {
	GetBookingByID(ctx context.Context, id string) (*dto.BookingResponse, error)
	GetBookingsByUserID(ctx context.Context, userID string) ([]dto.BookingResponse, error)
	CreateBooking(ctx context.Context, userID string, req dto.CreateBookingRequest) (*dto.BookingResponse, error)
	UpdateBookingStatus(ctx context.Context, id string, req dto.UpdateBookingStatusRequest) error
	DeleteBooking(ctx context.Context, id string) error
	CancelBooking(ctx context.Context, id string) error
}
