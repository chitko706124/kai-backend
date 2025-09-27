package domain

import (
	"context"

	"github.com/LouisFernando1204/kai-backend.git/dto"
)

type AuthService interface {
	Register(ctx context.Context, req dto.RegisterRequest) error
	Login(ctx context.Context, req dto.LoginRequest) (*dto.AuthResponse, error)
}
