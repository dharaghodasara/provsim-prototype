package main

import (
	"log"

	"github.com/elangovans/provsim-prototype/handlers"
	"github.com/elangovans/provsim-prototype/registry"
	"github.com/gin-gonic/gin"
)

//Router ...
var Router *gin.Engine

//StorageProvider
var storageProvider registry.StorageProvider

func loadData(sp registry.StorageProvider) {
	registry.RegistryMap, registry.RegisryLoaded = sp.Fetch()

	if registry.RegisryLoaded != nil {
		log.Panicf("Bootstrap failed with error %v", registry.RegisryLoaded) // TODO: Later try lazy loading and retries
	}
}

func setupEndpoints() {
	Router = gin.Default()
	v1 := Router.Group("/v1")

	v1.GET("/ping", handlers.GetPing)
	v1.POST("/providers", handlers.PostProvider)
	v1.GET("/providers/:id", handlers.GetProvider)
	v1.PUT("/providers/:id", handlers.UpdateProvider)

	Router.Run(":8080")
}

func bootstrap() {
	var lsf = registry.LocalFileSystemStorageProvider{FilePath: "/data/"}
	go loadData(lsf)

	setupEndpoints()
}

func main() {
	bootstrap()
}
