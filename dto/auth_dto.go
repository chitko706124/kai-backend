package dto

// swagger:model RegisterRequest
type RegisterRequest struct {
	// User's full name (minimum 3 characters)
	// required: true
	// example: John Doe
	FullName string `json:"full_name" validate:"required,min=3"`
	// User's email address
	// required: true
	// example: john.doe@example.com
	Email string `json:"email" validate:"required,email"`
	// User's phone number (minimum 10 characters)
	// required: true
	// example: 081234567890
	PhoneNumber string `json:"phone_number" validate:"required,min=10"`
	// User's identity number (exactly 16 characters)
	// required: true
	// example: 1234567890123456
	IdentityNumber string `json:"identity_number" validate:"required,min=16,max=16"`
	// User's password (minimum 6 characters)
	// required: true
	// example: password123
	Password string `json:"password" validate:"required,min=6"`
	// Password confirmation (must match password)
	// required: true
	// example: password123
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
}

// swagger:model LoginRequest
type LoginRequest struct {
	// User's email address
	// required: true
	// example: john.doe@example.com
	Email string `json:"email" validate:"required,email"`
	// User's password
	// required: true
	// example: password123
	Password string `json:"password" validate:"required"`
}

// swagger:model AuthResponse
type AuthResponse struct {
	// JWT authentication token
	// example: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
	Token string `json:"token"`
}

// swagger:model UserResponse
type UserResponse struct {
	// User's unique identifier
	// example: 507f1f77bcf86cd799439011
	ID string `json:"id"`
	// User's full name
	// example: John Doe
	FullName string `json:"full_name"`
	// User's email address
	// example: john.doe@example.com
	Email string `json:"email"`
	// User's phone number
	// example: 081234567890
	PhoneNumber string `json:"phone_number"`
	// User's identity number
	// example: 1234567890123456
	IdentityNumber string `json:"identity_number"`
}
