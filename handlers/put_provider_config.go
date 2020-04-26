package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UpdateProvider When this handler is triggered, get the payload form the context
// - Make sure the input unmarshals as a Provider object.
// - If it unmarshals then
// 	- Open the file and update the file data.
//
func UpdateProvider(c *gin.Context) {
	c.String(http.StatusAccepted, "got the changes. will update")
}
