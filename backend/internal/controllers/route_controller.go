package controllers

import (
	"net/http"

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
	routes, err := c.routeService.GetAllRoutes()
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to retrieve routes")
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Routes retrieved successfully", routes)
}

func (c *RouteController) CreateRoute(ctx *gin.Context) {
	var input struct {
		OriginCity      string `json:"OriginCity" binding:"required"`
		OriginPool      string `json:"OriginPool" binding:"required"`
		DestinationCity string `json:"DestinationCity" binding:"required"`
		DestinationPool string `json:"DestinationPool" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request payload")
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
