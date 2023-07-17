package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type healthCheckController struct {
}

func NewHealthCheckController() *healthCheckController {
	return &healthCheckController{}
}

// HealthCheck - healthcheck
// @Summary - healthcheck
// @Description - healthcheck
// @Tags - Health
// @Accept json
// @Produce json
// @Success 200 {string} string "OK"
// @Router /smart-chat/v1/health [get]
func (h *healthCheckController) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
