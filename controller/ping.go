package controller

import (
	"avalon/util"
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
)

// PingController used to check service
func PingController(c *gin.Context) {
	c.JSON(http.StatusOK, util.SuccessResponse())
}
