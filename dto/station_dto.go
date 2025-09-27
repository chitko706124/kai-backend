package dto

// swagger:model StationRequest
type StationRequest struct {
	// Station name
	// required: true
	// example: Gambir Station
	Name string `json:"name" validate:"required"`
	// City where the station is located
	// required: true
	// example: Jakarta
	City string `json:"city" validate:"required"`
	// Station code (2-4 characters)
	// required: true
	// example: GMB
	Code string `json:"code" validate:"required,min=2,max=4"`
}

// swagger:model StationResponse
type StationResponse struct {
	// Station unique identifier
	// example: 507f1f77bcf86cd799439011
	ID string `json:"id"`
	// Station name
	// example: Gambir Station
	Name string `json:"name"`
	// City where the station is located
	// example: Jakarta
	City string `json:"city"`
	// Station code
	// example: GMB
	Code string `json:"code"`
}
