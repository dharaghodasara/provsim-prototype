package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetPing ...
func GetPing(c *gin.Context) {
	c.JSON(http.StatusOK, `{"message":"pong"}`)
}
