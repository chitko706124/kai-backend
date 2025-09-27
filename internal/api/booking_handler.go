package api

import (
	"net/http"

	"github.com/LouisFernando1204/kai-backend.git/domain"
	"github.com/LouisFernando1204/kai-backend.git/dto"
	"github.com/LouisFernando1204/kai-backend.git/internal/util"
	"github.com/gofiber/fiber/v2"
)

type bookingHandler struct {
	bookingService domain.BookingService
}

func NewBookingHandler(bookingSvc domain.BookingService) *bookingHandler {
	return &bookingHandler{
		bookingService: bookingSvc,
	}
}

// GetMyBookings godoc
// @Summary      Get my bookings
// @Description  Retrieve all bookings belonging to the authenticated user
// @Tags         Bookings
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Success      200  {object}  dto.Response[[]dto.BookingResponse]  "Bookings retrieved successfully"
// @Failure      401  {object}  dto.Response[any]                   "Unauthorized"
// @Failure      500  {object}  dto.Response[any]                   "Internal server error"
// @Router       /bookings [get]
func (h *bookingHandler) GetMyBookings(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)

	bookings, err := h.bookingService.GetBookingsByUserID(c.Context(), userID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.CreateResponseWithoutData(http.StatusInternalServerError, err.Error()))
	}

	return c.Status(http.StatusOK).JSON(dto.CreateResponse(http.StatusOK, "Bookings retrieved successfully", bookings))
}

// GetBookingByID godoc
// @Summary      Get booking by ID
// @Description  Retrieve a specific booking by its ID
// @Tags         Bookings
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param        id   path      string                            true  "Booking ID"
// @Success      200  {object}  dto.Response[dto.BookingResponse]  "Booking retrieved successfully"
// @Failure      401  {object}  dto.Response[any]                 "Unauthorized"
// @Failure      500  {object}  dto.Response[any]                 "Internal server error"
// @Router       /bookings/{id} [get]
func (h *bookingHandler) GetBookingByID(c *fiber.Ctx) error {
	id := c.Params("id")
	booking, err := h.bookingService.GetBookingByID(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.CreateResponseWithoutData(http.StatusInternalServerError, err.Error()))
	}

	return c.Status(http.StatusOK).JSON(dto.CreateResponse(http.StatusOK, "Booking retrieved successfully", booking))
}

// CreateBooking godoc
// @Summary      Create a new booking
// @Description  Create a new train ticket booking for authenticated user
// @Tags         Bookings
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param        booking  body      dto.CreateBookingRequest      true  "Booking data"
// @Success      201      {object}  dto.Response[dto.BookingResponse]  "Booking created successfully"
// @Failure      400      {object}  dto.Response[any]             "Bad request - validation failed"
// @Failure      401      {object}  dto.Response[any]             "Unauthorized"
// @Failure      500      {object}  dto.Response[any]             "Internal server error"
// @Router       /bookings [post]
func (h *bookingHandler) CreateBooking(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)

	var req dto.CreateBookingRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.CreateResponseWithoutData(http.StatusBadRequest, err.Error()))
	}

	if errs := util.Validate(req); errs != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.CreateResponse(http.StatusBadRequest, "Validation failed", errs))
	}

	booking, err := h.bookingService.CreateBooking(c.Context(), userID, req)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.CreateResponseWithoutData(http.StatusInternalServerError, err.Error()))
	}

	return c.Status(http.StatusCreated).JSON(dto.CreateResponse(http.StatusCreated, "Booking created successfully", booking))
}

// UpdateStatus godoc
// @Summary      Update booking status
// @Description  Update the status of a booking (e.g., after payment)
// @Tags         Bookings
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param        id      path      string                            true  "Booking ID"
// @Param        status  body      dto.UpdateBookingStatusRequest    true  "New status data"
// @Success      200     {object}  dto.Response[any]                 "Booking status updated successfully"
// @Failure      400     {object}  dto.Response[any]                 "Bad request - validation failed"
// @Failure      401     {object}  dto.Response[any]                 "Unauthorized"
// @Failure      500     {object}  dto.Response[any]                 "Internal server error"
// @Router       /bookings/{id}/status [patch]
func (h *bookingHandler) UpdateStatus(c *fiber.Ctx) error {
	id := c.Params("id")
	var req dto.UpdateBookingStatusRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.CreateResponseWithoutData(http.StatusBadRequest, err.Error()))
	}

	if errs := util.Validate(req); errs != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.CreateResponse(http.StatusBadRequest, "Validation failed", errs))
	}

	err := h.bookingService.UpdateBookingStatus(c.Context(), id, req)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.CreateResponseWithoutData(http.StatusInternalServerError, err.Error()))
	}

	return c.Status(http.StatusOK).JSON(dto.CreateResponseWithoutData(http.StatusOK, "Booking status updated successfully"))
}

// CancelBooking godoc
// @Summary      Cancel booking
// @Description  Cancel an existing booking
// @Tags         Bookings
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param        id   path      string            true  "Booking ID"
// @Success      200  {object}  dto.Response[any]  "Booking cancelled successfully"
// @Failure      401  {object}  dto.Response[any]  "Unauthorized"
// @Failure      500  {object}  dto.Response[any]  "Internal server error"
// @Router       /bookings/{id}/cancel [post]
func (h *bookingHandler) CancelBooking(c *fiber.Ctx) error {
	id := c.Params("id")
	err := h.bookingService.CancelBooking(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.CreateResponseWithoutData(http.StatusInternalServerError, err.Error()))
	}

	return c.Status(http.StatusOK).JSON(dto.CreateResponseWithoutData(http.StatusOK, "Booking cancelled successfully"))
}
