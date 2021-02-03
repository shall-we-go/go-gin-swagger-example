package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheckHandler godoc
// @Summary Health Check
// @Description Health checking for the service
// @ID HealthCheckHandler
// @Produce plain
// @Success 200 {string} string "OK"
// @Router /healthcheck [get]
func HealthCheckHandler(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
