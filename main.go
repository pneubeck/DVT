package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pneubeck/DVT/controller"
)

func main() {
	router := gin.Default()
	router.GET("/vehicles", controller.GetVehicles)
	router.POST("/vehicles", controller.PostVehicle)

	router.Run("localhost:8080")
}
