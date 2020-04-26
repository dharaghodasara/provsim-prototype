package handlers

import (
	"fmt"
	"net/http"

	"github.com/elangovans/provsim-prototype/core"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gopkg.in/yaml.v2"
)

//PostProvider ...
func PostProvider(c *gin.Context) {
	fmt.Println("inside PostProvider")
	id, err := uuid.NewUUID()
	fmt.Println(id, err)
	p := core.Provider{
		Name: "A Provider Simulator",
		ID:   id,
	}

	s, err := yaml.Marshal(p)
	c.YAML(http.StatusCreated, &s)
}
