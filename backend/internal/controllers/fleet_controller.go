package controllers

import (
	"net/http"
	"strconv"

	"mesenshuttle-backend/internal/dto"
	"mesenshuttle-backend/internal/models"
	"mesenshuttle-backend/internal/services"
	"mesenshuttle-backend/pkg/utils"

	"github.com/gin-gonic/gin"
)

type FleetController struct {
	fleetService services.FleetService
}

func NewFleetController(fleetService services.FleetService) *FleetController {
	return &FleetController{fleetService}
}

func (c *FleetController) GetFleets(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	fleets, totalCount, err := c.fleetService.GetAllFleets(page, limit)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to retrieve fleets")
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Fleets retrieved successfully", dto.PaginatedResponse{
		Data:       fleets,
		TotalCount: totalCount,
		Page:       page,
		Limit:      limit,
	})
}

func (c *FleetController) CreateFleet(ctx *gin.Context) {
	var input dto.CreateFleetRequest

	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.ValidationErrorResponse(ctx, http.StatusBadRequest, utils.FormatValidationErrors(err))
		return
	}

	fleet := models.Fleet{
		Name:        input.Name,
		PlateNumber: input.PlateNumber,
		Type:        input.Type,
		TotalSeats:  input.TotalSeats,
	}

	if err := c.fleetService.CreateFleet(&fleet); err != nil {
		utils.HandleError(ctx, err)
		return
	}

	utils.SuccessResponse(ctx, http.StatusCreated, "Fleet created successfully", fleet)
}

func (c *FleetController) UpdateFleet(ctx *gin.Context) {
	id := ctx.Param("id")

	var input dto.UpdateFleetRequest
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.ValidationErrorResponse(ctx, http.StatusBadRequest, utils.FormatValidationErrors(err))
		return
	}

	fleet, err := c.fleetService.UpdateFleet(id, &input)
	if err != nil {
		utils.HandleError(ctx, err)
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Fleet updated successfully", fleet)
}
