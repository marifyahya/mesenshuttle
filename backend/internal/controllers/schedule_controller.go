package controllers

import (
	"net/http"
	"strconv"

	"mesenshuttle-backend/internal/dto"
	"mesenshuttle-backend/internal/services"
	"mesenshuttle-backend/pkg/utils"

	"github.com/gin-gonic/gin"
)

type ScheduleController struct {
	scheduleService services.ScheduleService
}

func NewScheduleController(scheduleService services.ScheduleService) *ScheduleController {
	return &ScheduleController{scheduleService}
}

func (c *ScheduleController) GetSchedules(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	schedules, totalCount, err := c.scheduleService.GetAllSchedules(page, limit)
	if err != nil {
		utils.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to retrieve schedules")
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Schedules retrieved successfully", dto.PaginatedResponse{
		Data:       schedules,
		TotalCount: totalCount,
		Page:       page,
		Limit:      limit,
	})
}
