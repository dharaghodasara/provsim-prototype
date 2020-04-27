package handlers

import (
	"log"
	"net/http"

	"github.com/elangovans/provsim-prototype/registry"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gopkg.in/yaml.v2"
)

//GetProvider ...
func GetProvider(c *gin.Context) {
	id := c.Params.ByName("id")
	uid := uuid.MustParse(id)
	p, found := registry.RegistryMap[uid]

	// TODO: Result handling can be optimized
	if !found {
		c.String(http.StatusNotFound, "Provider config not found with ID "+id)
	} else {
		log.Println(p)
		s, err := yaml.Marshal(&p)
		if err != nil {
			c.String(500, "Couldnt unmarshal Provider to Yaml")
		} else {
			c.YAML(http.StatusOK, string(s))
		}
	}
}
