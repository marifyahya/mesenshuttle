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

type RouteController struct {
	routeService services.RouteService
}

func NewRouteController(routeService services.RouteService) *RouteController {
	return &RouteController{routeService}
}

func (c *RouteController) GetRoutes(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	routes, totalCount, err := c.routeService.GetAllRoutes(page, limit)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to retrieve routes")
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Routes retrieved successfully", dto.PaginatedResponse{
		Data:       routes,
		TotalCount: totalCount,
		Page:       page,
		Limit:      limit,
	})
}

func (c *RouteController) CreateRoute(ctx *gin.Context) {
	var input dto.CreateRouteRequest

	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.ValidationErrorResponse(ctx, http.StatusBadRequest, utils.FormatValidationErrors(err))
		return
	}

	route := models.Route{
		OriginCity:      input.OriginCity,
		OriginPool:      input.OriginPool,
		DestinationCity: input.DestinationCity,
		DestinationPool: input.DestinationPool,
	}

	if err := c.routeService.CreateRoute(&route); err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to create route")
		return
	}

	utils.SuccessResponse(ctx, http.StatusCreated, "Route created successfully", route)
}
