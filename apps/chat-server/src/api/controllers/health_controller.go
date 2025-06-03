package controllers

import (
	"net/http"

	"github.com/shabashab/chattin/apps/chat-server/src/services"

	"github.com/gin-gonic/gin"
)

type HealthController struct {
	healthService *services.HealthService
}

func NewHealthController(healthService *services.HealthService) *HealthController {
	return &HealthController{
		healthService: healthService,
	}
}

// GetHealth godoc
// @Summary Get health status
// @Description Get the health status of the server
// @Tags health
// @Accept  json
// @Produce  json
// @Success 200 {object} services.HealthStatus
// @Router /health [get]
func (controller HealthController) GetHealth(ctx *gin.Context) {
	healthStatus := controller.healthService.GetHealthStatus()
	httpStatus := http.StatusOK

	if !healthStatus.Alive {
		httpStatus = http.StatusInternalServerError
	}

	if httpStatus == http.StatusInternalServerError {
		healthStatus.Alive = false
	}

	ctx.JSON(httpStatus, gin.H{
		"status": healthStatus,
	})
}
