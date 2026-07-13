package controllers

import (
	"net/http"

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
