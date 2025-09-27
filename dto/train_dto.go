package dto

// swagger:model SeatDTO
type SeatDTO struct {
	// Seat code (e.g., A1, B2)
	// required: true
	// example: A1
	Code string `json:"code" validate:"required"`
	// Whether the seat is available for booking
	// example: true
	IsAvailable bool `json:"is_available"`
}

// swagger:model CarriageDTO
type CarriageDTO struct {
	// Carriage code (e.g., CAR1, CAR2)
	// required: true
	// example: CAR1
	Code string `json:"code" validate:"required"`
	// List of seats in this carriage
	// required: true
	Seats []SeatDTO `json:"seats" validate:"required"`
}

// swagger:model CreateTrainRequest
type CreateTrainRequest struct {
	// Train name
	// required: true
	// example: Argo Bromo Anggrek
	Name string `json:"name" validate:"required"`
	// Train class (e.g., Executive, Business, Economy)
	// required: true
	// example: Executive
	Class string `json:"class" validate:"required"`
	// List of carriages with their seats
	// required: true
	Carriages []CarriageDTO `json:"carriages" validate:"required"`
}

// swagger:model UpdateTrainRequest
type UpdateTrainRequest struct {
	// Train name
	// required: true
	// example: Argo Bromo Anggrek
	Name string `json:"name" validate:"required"`
	// Train class (e.g., Executive, Business, Economy)
	// required: true
	// example: Executive
	Class string `json:"class" validate:"required"`
	// List of carriages with their seats
	// required: true
	Carriages []CarriageDTO `json:"carriages" validate:"required"`
}

// swagger:model TrainResponse
type TrainResponse struct {
	// Train unique identifier
	// example: 507f1f77bcf86cd799439011
	ID string `json:"id"`
	// Train name
	// example: Argo Bromo Anggrek
	Name string `json:"name"`
	// Train class
	// example: Executive
	Class string `json:"class"`
	// List of carriages with their seats
	Carriages []CarriageDTO `json:"carriages"`
}
