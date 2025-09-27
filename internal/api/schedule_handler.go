package api

import (
	"net/http"

	"github.com/LouisFernando1204/kai-backend.git/domain"
	"github.com/LouisFernando1204/kai-backend.git/dto"
	"github.com/LouisFernando1204/kai-backend.git/internal/util"
	"github.com/gofiber/fiber/v2"
)

type scheduleHandler struct {
	scheduleService domain.ScheduleService
}

func NewScheduleHandler(scheduleSvc domain.ScheduleService) *scheduleHandler {
	return &scheduleHandler{
		scheduleService: scheduleSvc,
	}
}

// GetAll godoc
// @Summary      Get all schedules (Admin only)
// @Description  Retrieve all train schedules for admin management
// @Tags         Schedules
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Success      200  {object}  dto.Response[[]dto.ScheduleResponse]  "All schedules retrieved successfully"
// @Failure      401  {object}  dto.Response[any]                    "Unauthorized"
// @Failure      500  {object}  dto.Response[any]                    "Internal server error"
// @Router       /schedules [get]
func (h *scheduleHandler) GetAll(c *fiber.Ctx) error {
	schedules, err := h.scheduleService.GetAll(c.Context())
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.CreateResponseWithoutData(http.StatusInternalServerError, err.Error()))
	}

	return c.Status(http.StatusOK).JSON(dto.CreateResponse(http.StatusOK, "All schedules retrieved successfully", schedules))
}

// GetByID godoc
// @Summary      Get schedule by ID
// @Description  Retrieve a specific train schedule by its ID
// @Tags         Schedules
// @Accept       json
// @Produce      json
// @Param        id   path      string                             true  "Schedule ID"
// @Success      200  {object}  dto.Response[dto.ScheduleResponse]  "Schedule retrieved successfully"
// @Failure      404  {object}  dto.Response[any]                  "Schedule not found"
// @Failure      500  {object}  dto.Response[any]                  "Internal server error"
// @Router       /schedules/{id} [get]
func (h *scheduleHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	schedule, err := h.scheduleService.GetByID(c.Context(), id)
	if err != nil {
		if err.Error() == "schedule not found" {
			return c.Status(http.StatusNotFound).JSON(dto.CreateResponseWithoutData(http.StatusNotFound, err.Error()))
		}
		return c.Status(http.StatusInternalServerError).JSON(dto.CreateResponseWithoutData(http.StatusInternalServerError, err.Error()))
	}

	return c.Status(http.StatusOK).JSON(dto.CreateResponse(http.StatusOK, "Schedule retrieved successfully", schedule))
}

// GetSeatLayout godoc
// @Summary      Get seat layout
// @Description  Retrieve the seat layout for a specific schedule
// @Tags         Schedules
// @Accept       json
// @Produce      json
// @Param        id   path      string                                     true  "Schedule ID"
// @Success      200  {object}  dto.Response[[]dto.CarriageLayoutDTO]      "Seat layout retrieved successfully"
// @Failure      500  {object}  dto.Response[any]                          "Internal server error"
// @Router       /schedules/{id}/seats [get]
func (h *scheduleHandler) GetSeatLayout(c *fiber.Ctx) error {
	id := c.Params("id")
	layout, err := h.scheduleService.GetSeatLayout(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.CreateResponseWithoutData(http.StatusInternalServerError, err.Error()))
	}

	return c.Status(http.StatusOK).JSON(dto.CreateResponse(http.StatusOK, "Seat layout retrieved successfully", layout))
}

// Create godoc
// @Summary      Create a new schedule
// @Description  Create a new train schedule (Admin only)
// @Tags         Schedules
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param        schedule  body      dto.CreateScheduleRequest        true  "Schedule data"
// @Success      201       {object}  dto.Response[dto.ScheduleResponse]  "Schedule created successfully"
// @Failure      400       {object}  dto.Response[any]                "Bad request - validation failed"
// @Failure      401       {object}  dto.Response[any]                "Unauthorized"
// @Failure      500       {object}  dto.Response[any]                "Internal server error"
// @Router       /schedules [post]
func (h *scheduleHandler) Create(c *fiber.Ctx) error {
	var req dto.CreateScheduleRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.CreateResponseWithoutData(http.StatusBadRequest, err.Error()))
	}

	if errs := util.Validate(req); errs != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.CreateResponse(http.StatusBadRequest, "Validation failed", errs))
	}

	schedule, err := h.scheduleService.Create(c.Context(), req)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.CreateResponseWithoutData(http.StatusInternalServerError, err.Error()))
	}

	return c.Status(http.StatusCreated).JSON(dto.CreateResponse(http.StatusCreated, "Schedule created successfully", schedule))
}

// Update godoc
// @Summary      Update schedule
// @Description  Update an existing train schedule (Admin only)
// @Tags         Schedules
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param        id        path      string                             true  "Schedule ID"
// @Param        schedule  body      dto.UpdateScheduleRequest          true  "Updated schedule data"
// @Success      200       {object}  dto.Response[dto.ScheduleResponse]  "Schedule updated successfully"
// @Failure      400       {object}  dto.Response[any]                  "Bad request - validation failed"
// @Failure      401       {object}  dto.Response[any]                  "Unauthorized"
// @Failure      500       {object}  dto.Response[any]                  "Internal server error"
// @Router       /schedules/{id} [put]
func (h *scheduleHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var req dto.UpdateScheduleRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.CreateResponseWithoutData(http.StatusBadRequest, err.Error()))
	}

	if errs := util.Validate(req); errs != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.CreateResponse(http.StatusBadRequest, "Validation failed", errs))
	}

	schedule, err := h.scheduleService.Update(c.Context(), id, req)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.CreateResponseWithoutData(http.StatusInternalServerError, err.Error()))
	}

	return c.Status(http.StatusOK).JSON(dto.CreateResponse(http.StatusOK, "Schedule updated successfully", schedule))
}

// Delete godoc
// @Summary      Delete schedule
// @Description  Delete a train schedule (Admin only)
// @Tags         Schedules
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param        id   path      string            true  "Schedule ID"
// @Success      200  {object}  dto.Response[any]  "Schedule deleted successfully"
// @Failure      401  {object}  dto.Response[any]  "Unauthorized"
// @Failure      500  {object}  dto.Response[any]  "Internal server error"
// @Router       /schedules/{id} [delete]
func (h *scheduleHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	err := h.scheduleService.Delete(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.CreateResponseWithoutData(http.StatusInternalServerError, err.Error()))
	}

	return c.Status(http.StatusOK).JSON(dto.CreateResponseWithoutData(http.StatusOK, "Schedule deleted successfully"))
}

// Search godoc
// @Summary      Search schedules
// @Description  Search train schedules by origin, destination, date, and passenger count
// @Tags         Schedules
// @Accept       json
// @Produce      json
// @Param        schedule  body      dto.SearchScheduleRequest        true  "Schedule data"
// @Success      200                     {object}  dto.Response[[]dto.ScheduleResponse]  "Schedules retrieved successfully"
// @Failure      400                     {object}  dto.Response[any]                    "Bad request - validation failed"
// @Failure      500                     {object}  dto.Response[any]                    "Internal server error"
// @Router       /schedules/search [post]
func (h *scheduleHandler) Search(c *fiber.Ctx) error {
	var req dto.SearchScheduleRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.CreateResponseWithoutData(http.StatusBadRequest, err.Error()))
	}

	if errs := util.Validate(req); errs != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.CreateResponse(http.StatusBadRequest, "Validation failed", errs))
	}

	schedules, err := h.scheduleService.Search(c.Context(), req)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.CreateResponseWithoutData(http.StatusInternalServerError, err.Error()))
	}

	return c.Status(http.StatusOK).JSON(dto.CreateResponse(http.StatusOK, "Schedules retrieved successfully", schedules))
}
