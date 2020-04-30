package main

import (
	"log"

	"provsim-prototype/handlers"
	"provsim-prototype/registry"

	"github.com/gin-gonic/gin"
)

//Router ...
var Router *gin.Engine

func loadData(sp registry.StorageProvider) {
	registry.RegistryMap, registry.RegisryLoaded = sp.Fetch()

	if registry.RegisryLoaded != nil {
		log.Panicf("Bootstrap failed with error %v", registry.RegisryLoaded) // TODO: Later try lazy loading and retries
	}
}

func setupEndpoints() {
	Router = gin.Default()
	//gin.SetMode(gin.ReleaseMode)
	v1 := Router.Group("/v1")

	v1.GET("/ping", handlers.GetPing)
	v1.GET("/providers/:id", handlers.GetProvider)
	v1.PUT("/providers/:id", handlers.UpdateProvider)
	v1.POST("/providers", handlers.PostProvider)

	Router.Run(":8080")
}

func bootstrap() {
	var lsf = registry.LocalFileSystemStorageProvider{
		FilePath: "/data/",
	}
	go loadData(lsf) //load the configuration data concurrently (async)
	setupEndpoints()
}

func main() {
	bootstrap()
}
