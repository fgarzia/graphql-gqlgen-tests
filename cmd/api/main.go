package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/go-meli-toolkit/gingonic/mlhandlers"
)

func main() {
	if err := SetupRouter().Run(":8080"); err != nil {
		panic(err.Error())
	}
}

func SetupRouter() *gin.Engine {
	router := mlhandlers.DefaultMeliRouter()

	mapRoutes(router)

	return router
}
