package api

import (
	"net/http"

	"github.com/LouisFernando1204/kai-backend.git/domain"
	"github.com/LouisFernando1204/kai-backend.git/dto"
	"github.com/LouisFernando1204/kai-backend.git/internal/util"
	"github.com/gofiber/fiber/v2"
)

type trainHandler struct {
	trainService domain.TrainService
}

func NewTrainHandler(trainSvc domain.TrainService) *trainHandler {
	return &trainHandler{
		trainService: trainSvc,
	}
}

// GetAll godoc
// @Summary      Get all trains
// @Description  Retrieve a list of all trains
// @Tags         Trains
// @Accept       json
// @Produce      json
// @Success      200  {object}  dto.Response[[]dto.TrainResponse]  "Trains retrieved successfully"
// @Failure      500  {object}  dto.Response[any]                 "Internal server error"
// @Router       /trains [get]
func (h *trainHandler) GetAll(c *fiber.Ctx) error {
	trains, err := h.trainService.GetAll(c.Context())
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.CreateResponseWithoutData(http.StatusInternalServerError, err.Error()))
	}

	return c.Status(http.StatusOK).JSON(dto.CreateResponse(http.StatusOK, "Trains retrieved successfully", trains))
}

// GetByID godoc
// @Summary      Get train by ID
// @Description  Retrieve a specific train by its ID
// @Tags         Trains
// @Accept       json
// @Produce      json
// @Param        id   path      string                         true  "Train ID"
// @Success      200  {object}  dto.Response[dto.TrainResponse]  "Train retrieved successfully"
// @Failure      404  {object}  dto.Response[any]              "Train not found"
// @Failure      500  {object}  dto.Response[any]              "Internal server error"
// @Router       /trains/{id} [get]
func (h *trainHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	train, err := h.trainService.GetByID(c.Context(), id)
	if err != nil {
		if err.Error() == "train not found" {
			return c.Status(http.StatusNotFound).JSON(dto.CreateResponseWithoutData(http.StatusNotFound, err.Error()))
		}
		return c.Status(http.StatusInternalServerError).JSON(dto.CreateResponseWithoutData(http.StatusInternalServerError, err.Error()))
	}

	return c.Status(http.StatusOK).JSON(dto.CreateResponse(http.StatusOK, "Train retrieved successfully", train))
}

// Create godoc
// @Summary      Create a new train
// @Description  Create a new train with carriages and seats (Admin only)
// @Tags         Trains
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param        train  body      dto.CreateTrainRequest       true  "Train data"
// @Success      201    {object}  dto.Response[dto.TrainResponse]  "Train created successfully"
// @Failure      400    {object}  dto.Response[any]            "Bad request - validation failed"
// @Failure      401    {object}  dto.Response[any]            "Unauthorized"
// @Failure      500    {object}  dto.Response[any]            "Internal server error"
// @Router       /trains [post]
func (h *trainHandler) Create(c *fiber.Ctx) error {
	var req dto.CreateTrainRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.CreateResponseWithoutData(http.StatusBadRequest, err.Error()))
	}

	if errs := util.Validate(req); errs != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.CreateResponse(http.StatusBadRequest, "Validation failed", errs))
	}

	train, err := h.trainService.Create(c.Context(), req)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.CreateResponseWithoutData(http.StatusInternalServerError, err.Error()))
	}

	return c.Status(http.StatusCreated).JSON(dto.CreateResponse(http.StatusCreated, "Train created successfully", train))
}

// Update godoc
// @Summary      Update train
// @Description  Update an existing train (Admin only)
// @Tags         Trains
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param        id     path      string                         true  "Train ID"
// @Param        train  body      dto.UpdateTrainRequest         true  "Updated train data"
// @Success      200    {object}  dto.Response[dto.TrainResponse]  "Train updated successfully"
// @Failure      400    {object}  dto.Response[any]              "Bad request - validation failed"
// @Failure      401    {object}  dto.Response[any]              "Unauthorized"
// @Failure      500    {object}  dto.Response[any]              "Internal server error"
// @Router       /trains/{id} [put]
func (h *trainHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var req dto.UpdateTrainRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.CreateResponseWithoutData(http.StatusBadRequest, err.Error()))
	}

	if errs := util.Validate(req); errs != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.CreateResponse(http.StatusBadRequest, "Validation failed", errs))
	}

	train, err := h.trainService.Update(c.Context(), id, req)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.CreateResponseWithoutData(http.StatusInternalServerError, err.Error()))
	}

	return c.Status(http.StatusOK).JSON(dto.CreateResponse(http.StatusOK, "Train updated successfully", train))
}

// Delete godoc
// @Summary      Delete train
// @Description  Delete a train (Admin only)
// @Tags         Trains
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param        id   path      string            true  "Train ID"
// @Success      200  {object}  dto.Response[any]  "Train deleted successfully"
// @Failure      401  {object}  dto.Response[any]  "Unauthorized"
// @Failure      500  {object}  dto.Response[any]  "Internal server error"
// @Router       /trains/{id} [delete]
func (h *trainHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	err := h.trainService.Delete(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.CreateResponseWithoutData(http.StatusInternalServerError, err.Error()))
	}

	return c.Status(http.StatusOK).JSON(dto.CreateResponseWithoutData(http.StatusOK, "Train deleted successfully"))
}
