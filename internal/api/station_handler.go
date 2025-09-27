package api

import (
	"net/http"

	"github.com/LouisFernando1204/kai-backend.git/domain"
	"github.com/LouisFernando1204/kai-backend.git/dto"
	"github.com/LouisFernando1204/kai-backend.git/internal/util"
	"github.com/gofiber/fiber/v2"
)

type stationHandler struct {
	stationService domain.StationService
}

func NewStationHandler(stationSvc domain.StationService) *stationHandler {
	return &stationHandler{
		stationService: stationSvc,
	}
}

// GetAll godoc
// @Summary      Get all stations
// @Description  Retrieve a list of all railway stations
// @Tags         Stations
// @Accept       json
// @Produce      json
// @Success      200  {object}  dto.Response[[]dto.StationResponse]  "Stations retrieved successfully"
// @Failure      500  {object}  dto.Response[any]                   "Internal server error"
// @Router       /stations [get]
func (h *stationHandler) GetAll(c *fiber.Ctx) error {
	stations, err := h.stationService.GetAll(c.Context())
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.CreateResponseWithoutData(http.StatusInternalServerError, err.Error()))
	}

	return c.Status(http.StatusOK).JSON(dto.CreateResponse(http.StatusOK, "Stations retrieved successfully", stations))
}

// GetByID godoc
// @Summary      Get station by ID
// @Description  Retrieve a specific station by its ID
// @Tags         Stations
// @Accept       json
// @Produce      json
// @Param        id   path      string                           true  "Station ID"
// @Success      200  {object}  dto.Response[dto.StationResponse]  "Station retrieved successfully"
// @Failure      404  {object}  dto.Response[any]                "Station not found"
// @Failure      500  {object}  dto.Response[any]                "Internal server error"
// @Router       /stations/{id} [get]
func (h *stationHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	station, err := h.stationService.GetByID(c.Context(), id)
	if err != nil {
		if err.Error() == "station not found" {
			return c.Status(http.StatusNotFound).JSON(dto.CreateResponseWithoutData(http.StatusNotFound, err.Error()))
		}
		return c.Status(http.StatusInternalServerError).JSON(dto.CreateResponseWithoutData(http.StatusInternalServerError, err.Error()))
	}

	return c.Status(http.StatusOK).JSON(dto.CreateResponse(http.StatusOK, "Station retrieved successfully", station))
}

// Create godoc
// @Summary      Create a new station
// @Description  Create a new railway station (Admin only)
// @Tags         Stations
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param        station  body      dto.StationRequest           true  "Station data"
// @Success      201      {object}  dto.Response[dto.StationResponse]  "Station created successfully"
// @Failure      400      {object}  dto.Response[any]            "Bad request - validation failed"
// @Failure      401      {object}  dto.Response[any]            "Unauthorized"
// @Failure      500      {object}  dto.Response[any]            "Internal server error"
// @Router       /stations [post]
func (h *stationHandler) Create(c *fiber.Ctx) error {
	var req dto.StationRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.CreateResponseWithoutData(http.StatusBadRequest, err.Error()))
	}

	if errs := util.Validate(req); errs != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.CreateResponse(http.StatusBadRequest, "Validation failed", errs))
	}

	station, err := h.stationService.Create(c.Context(), req)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.CreateResponseWithoutData(http.StatusInternalServerError, err.Error()))
	}

	return c.Status(http.StatusCreated).JSON(dto.CreateResponse(http.StatusCreated, "Station created successfully", station))
}

// Update godoc
// @Summary      Update station
// @Description  Update an existing railway station (Admin only)
// @Tags         Stations
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param        id       path      string                           true  "Station ID"
// @Param        station  body      dto.StationRequest               true  "Updated station data"
// @Success      200      {object}  dto.Response[dto.StationResponse]  "Station updated successfully"
// @Failure      400      {object}  dto.Response[any]                "Bad request - validation failed"
// @Failure      401      {object}  dto.Response[any]                "Unauthorized"
// @Failure      500      {object}  dto.Response[any]                "Internal server error"
// @Router       /stations/{id} [put]
func (h *stationHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var req dto.StationRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.CreateResponseWithoutData(http.StatusBadRequest, err.Error()))
	}

	if errs := util.Validate(req); errs != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.CreateResponse(http.StatusBadRequest, "Validation failed", errs))
	}

	station, err := h.stationService.Update(c.Context(), id, req)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.CreateResponseWithoutData(http.StatusInternalServerError, err.Error()))
	}

	return c.Status(http.StatusOK).JSON(dto.CreateResponse(http.StatusOK, "Station updated successfully", station))
}

// Delete godoc
// @Summary      Delete station
// @Description  Delete a railway station (Admin only)
// @Tags         Stations
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param        id   path      string            true  "Station ID"
// @Success      200  {object}  dto.Response[any]  "Station deleted successfully"
// @Failure      401  {object}  dto.Response[any]  "Unauthorized"
// @Failure      500  {object}  dto.Response[any]  "Internal server error"
// @Router       /stations/{id} [delete]
func (h *stationHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	err := h.stationService.Delete(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.CreateResponseWithoutData(http.StatusInternalServerError, err.Error()))
	}

	return c.Status(http.StatusOK).JSON(dto.CreateResponseWithoutData(http.StatusOK, "Station deleted successfully"))
}
