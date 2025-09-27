package dto

// swagger:model SearchScheduleRequest
type SearchScheduleRequest struct {
	// Origin station ID
	// required: true
	// example: 507f1f77bcf86cd799439011
	OriginStationID string `json:"origin_station_id" validate:"required"`
	// Destination station ID
	// required: true
	// example: 507f1f77bcf86cd799439012
	DestinationStationID string `json:"destination_station_id" validate:"required"`
	// Departure date in YYYY-MM-DD format
	// required: true
	// example: 2024-01-15
	DepartureDate string `json:"departure_date" validate:"required"`
	// Number of passengers
	// required: true
	// minimum: 1
	// example: 2
	Passengers int `json:"passengers" validate:"required,min=1"`
}

// swagger:model CreateScheduleRequest
type CreateScheduleRequest struct {
	// Train ID
	// required: true
	// example: 507f1f77bcf86cd799439013
	TrainID string `json:"train_id" validate:"required"`
	// Origin station ID
	// required: true
	// example: 507f1f77bcf86cd799439011
	OriginStationID string `json:"origin_station_id" validate:"required"`
	// Destination station ID
	// required: true
	// example: 507f1f77bcf86cd799439012
	DestinationStationID string `json:"destination_station_id" validate:"required"`
	// Departure time in YYYY-MM-DD HH:MM format
	// required: true
	// example: 2024-01-15 08:30
	DepartureTime string `json:"departure_time" validate:"required"`
	// Arrival time in YYYY-MM-DD HH:MM format
	// required: true
	// example: 2024-01-15 16:45
	ArrivalTime string `json:"arrival_time" validate:"required"`
	// Ticket price
	// required: true
	// minimum: 0
	// example: 150000
	Price float64 `json:"price" validate:"required,gt=0"`
	// Number of available seats
	// required: true
	// minimum: 0
	// example: 50
	AvailableSeats int `json:"available_seats" validate:"required,gte=0"`
}

// swagger:model UpdateScheduleRequest
type UpdateScheduleRequest struct {
	// Departure time in YYYY-MM-DD HH:MM format
	// required: true
	// example: 2024-01-15 08:30
	DepartureTime string `json:"departure_time" validate:"required"`
	// Arrival time in YYYY-MM-DD HH:MM format
	// required: true
	// example: 2024-01-15 16:45
	ArrivalTime string `json:"arrival_time" validate:"required"`
	// Ticket price
	// required: true
	// minimum: 0
	// example: 150000
	Price float64 `json:"price" validate:"required,gt=0"`
	// Number of available seats
	// required: true
	// minimum: 0
	// example: 50
	AvailableSeats int `json:"available_seats" validate:"required,gte=0"`
}

// swagger:model SeatAvailabilityDTO
type SeatAvailabilityDTO struct {
	// Seat code (e.g., A1, B2)
	// example: A1
	Code string `json:"code"`
	// Whether the seat is available for booking
	// example: true
	IsAvailable bool `json:"is_available"`
}

// swagger:model CarriageLayoutDTO
type CarriageLayoutDTO struct {
	// Carriage code
	// example: CAR1
	Code string `json:"code"`
	// List of seats with their availability
	Seats []SeatAvailabilityDTO `json:"seats"`
}

// swagger:model TrainInScheduleResponse
type TrainInScheduleResponse struct {
	// Train name
	// example: Argo Bromo Anggrek
	Name string `json:"name"`
	// Train class
	// example: Executive
	Class string `json:"class"`
}

// swagger:model ScheduleResponse
type ScheduleResponse struct {
	// Schedule unique identifier
	// example: 507f1f77bcf86cd799439014
	ID string `json:"id"`
	// Train information
	Train TrainInScheduleResponse `json:"train"`
	// Origin station information
	OriginStation StationResponse `json:"origin_station"`
	// Destination station information
	DestinationStation StationResponse `json:"destination_station"`
	// Departure time in YYYY-MM-DD HH:MM format
	// example: 2024-01-15 08:30
	DepartureTime string `json:"departure_time"`
	// Arrival time in YYYY-MM-DD HH:MM format
	// example: 2024-01-15 16:45
	ArrivalTime string `json:"arrival_time"`
	// Journey duration (e.g., "8h 15m")
	// example: 8h 15m
	Duration string `json:"duration"`
	// Ticket price
	// example: 150000
	Price float64 `json:"price"`
	// Number of available seats
	// example: 25
	AvailableSeats int `json:"available_seats"`
}
