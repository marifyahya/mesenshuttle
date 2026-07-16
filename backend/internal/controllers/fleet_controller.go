package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"mesenshuttle-backend/internal/dto"
	"mesenshuttle-backend/internal/models"
	"mesenshuttle-backend/internal/services"
	"mesenshuttle-backend/pkg/apperrors"
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
		// Primary guard: sentinel error from Service layer
		if err == apperrors.ErrDuplicatePlateNumber {
			utils.ValidationErrorResponse(ctx, http.StatusBadRequest, []dto.FieldError{
				{Field: "plate_number", Message: "plate number is already registered"},
			})
			return
		}
		// Fallback guard: race condition caused duplicate entry at DB level
		if strings.Contains(err.Error(), "Duplicate entry") && strings.Contains(err.Error(), "plate_number") {
			utils.ValidationErrorResponse(ctx, http.StatusBadRequest, []dto.FieldError{
				{Field: "plate_number", Message: "plate number is already registered"},
			})
			return
		}
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to create fleet")
		return
	}

	utils.SuccessResponse(ctx, http.StatusCreated, "Fleet created successfully", fleet)
}
