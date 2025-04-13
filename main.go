package main

import (
	"github.com/gin-gonic/gin"
	VehicleController "github.com/pneubeck/DVT/controller"
)

func main() {
	router := gin.Default()
	router.GET("/vehicles", VehicleController.GetVehicles)
	router.GET("/vehicles/:id", VehicleController.GetVehicleById)
	router.POST("/vehicles", VehicleController.PostVehicle)

	router.Run("localhost:8080")
}
