package dto

// swagger:model BookerRequest
type BookerRequest struct {
	// Booker's full name
	// required: true
	// example: John Doe
	FullName string `json:"full_name" validate:"required"`
	// Booker's email address
	// required: true
	// example: john.doe@example.com
	Email string `json:"email" validate:"required,email"`
	// Booker's phone number
	// required: true
	// example: 081234567890
	PhoneNumber string `json:"phone_number" validate:"required"`
}

// swagger:model PassengerRequest
type PassengerRequest struct {
	// Passenger's full name
	// required: true
	// example: Jane Doe
	FullName string `json:"full_name" validate:"required"`
	// Passenger's identity number (exactly 16 characters)
	// required: true
	// example: 1234567890123456
	IdentityNumber string `json:"identity_number" validate:"required,min=16,max=16"`
	// Seat code assigned to the passenger
	// required: true
	// example: EKO-6 / 9A
	SeatCode string `json:"seat_code" validate:"required"`
}

// swagger:model CreateBookingRequest
type CreateBookingRequest struct {
	// Schedule ID for the booking
	// required: true
	// example: 507f1f77bcf86cd799439014
	ScheduleID string `json:"schedule_id" validate:"required"`
	// Booker information
	// required: true
	Booker BookerRequest `json:"booker" validate:"required"`
	// List of passengers (minimum 1)
	// required: true
	Passengers []PassengerRequest `json:"passengers" validate:"required,min=1"`
}

// swagger:model UpdateBookingStatusRequest
type UpdateBookingStatusRequest struct {
	// New booking status (e.g., "confirmed", "paid", "cancelled")
	// required: true
	// example: paid
	Status string `json:"status" validate:"required"`
}

// swagger:model BookingResponse
type BookingResponse struct {
	// Booking unique identifier
	// example: 507f1f77bcf86cd799439015
	ID string `json:"id"`
	// Unique booking code
	// example: KAI-20240115-001
	BookingCode string `json:"booking_code"`
	// Schedule information
	Schedule ScheduleResponse `json:"schedule"`
	// List of passengers
	Passengers []PassengerRequest `json:"passengers"`
	// Total price for all passengers
	// example: 300000
	TotalPrice float64 `json:"total_price"`
	// Booking status
	// example: pending
	Status string `json:"status"`
	// Payment expiry time in YYYY-MM-DD HH:MM format
	// example: 2024-01-15 12:00
	PaymentExpiry string `json:"payment_expiry"`
}
