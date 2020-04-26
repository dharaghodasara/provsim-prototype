package handlers

import (
	"fmt"
	"net/http"

	"github.com/elangovans/provsim-prototype/data"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//GetProvider ...
func GetProvider(c *gin.Context) {
	id := c.Params.ByName("id")
	uuid := uuid.MustParse(id)

	fmt.Println(len(data.ConfigMap))

	p := data.ConfigMap[uuid]
	c.YAML(http.StatusOK, &p)
}
