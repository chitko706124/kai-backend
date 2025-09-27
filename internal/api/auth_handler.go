package api

import (
	"net/http"

	"github.com/LouisFernando1204/kai-backend.git/domain"
	"github.com/LouisFernando1204/kai-backend.git/dto"
	"github.com/LouisFernando1204/kai-backend.git/internal/util"
	"github.com/gofiber/fiber/v2"
)

type authHandler struct {
	authService domain.AuthService
}

func NewAuthHandler(authSvc domain.AuthService) *authHandler {
	return &authHandler{
		authService: authSvc,
	}
}

// Register godoc
// @Summary      Register a new user
// @Description  Create a new user account with the provided information
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Param        user  body      dto.RegisterRequest  true  "User registration data"
// @Success      201   {object}  dto.Response[any]    "User registered successfully"
// @Failure      400   {object}  dto.Response[any]    "Bad request - validation failed"
// @Failure      500   {object}  dto.Response[any]    "Internal server error"
// @Router       /auth/register [post]
func (h *authHandler) Register(c *fiber.Ctx) error {
	var req dto.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.CreateResponseWithoutData(http.StatusBadRequest, err.Error()))
	}

	if errs := util.Validate(req); errs != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.CreateResponse(http.StatusBadRequest, "Validation failed", errs))
	}

	err := h.authService.Register(c.Context(), req)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.CreateResponseWithoutData(http.StatusInternalServerError, err.Error()))
	}

	return c.Status(http.StatusCreated).JSON(dto.CreateResponseWithoutData(http.StatusCreated, "User registered successfully"))
}

// Login godoc
// @Summary      User login
// @Description  Authenticate user and return JWT token
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Param        credentials  body      dto.LoginRequest         true  "User login credentials"
// @Success      200          {object}  dto.Response[dto.AuthResponse]  "Login successful"
// @Failure      400          {object}  dto.Response[any]        "Bad request - validation failed"
// @Failure      401          {object}  dto.Response[any]        "Unauthorized - invalid credentials"
// @Router       /auth/login [post]
func (h *authHandler) Login(c *fiber.Ctx) error {
	var req dto.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.CreateResponseWithoutData(http.StatusBadRequest, err.Error()))
	}

	if errs := util.Validate(req); errs != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.CreateResponse(http.StatusBadRequest, "Validation failed", errs))
	}

	res, err := h.authService.Login(c.Context(), req)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(dto.CreateResponseWithoutData(http.StatusUnauthorized, err.Error()))
	}

	return c.Status(http.StatusOK).JSON(dto.CreateResponse(http.StatusOK, "Login successful", res))
}
