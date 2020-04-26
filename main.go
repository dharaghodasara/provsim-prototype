package main

import (
	"github.com/elangovans/provsim-prototype/data"
	"github.com/elangovans/provsim-prototype/handlers"
	"github.com/gin-gonic/gin"
)

//Router ...
var Router *gin.Engine

func bootstrap() {
	go data.LoadSimulationConfigurations()
	setupEndpoints()
}

func setupEndpoints() {
	Router = gin.Default()
	v1 := Router.Group("/v1")

	v1.GET("/ping", handlers.GetPing)
	v1.POST("/providers", handlers.PostProvider)
	v1.GET("/providers/:id", handlers.GetProvider)

	Router.Run(":8080")
}

func main() {
	bootstrap()
}
