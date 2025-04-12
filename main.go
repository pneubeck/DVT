package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pneubeck/DVT/controller"
)

func main() {
	router := gin.Default()
	router.GET("/vehicles", controller.GetVehicles)

	router.Run("localhost:8080")
}
