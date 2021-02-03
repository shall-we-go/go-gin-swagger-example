package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheckHandler godoc
// @summary Health Check
// @description Health checking for the service
// @id HealthCheckHandler
// @produce plain
// @response 200 {string} string "OK"
// @router /healthcheck [get]
func HealthCheckHandler(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
