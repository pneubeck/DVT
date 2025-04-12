package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pneubeck/DVT/structs"
)

// declaration of test vehicle data
var vehicles = []structs.Vehicle{
	{ID: 1, VehicleNo: "54365453242345", VehicleColor: "Blue"},
	{ID: 2, VehicleNo: "63243243432662", VehicleColor: "Yellow"},
	{ID: 3, VehicleNo: "43435766321454", VehicleColor: "Green"},
}

func main() {
	router := gin.Default()
	router.GET("/vehicles", getVehicles)

	router.Run("localhost:8080")
}

func getVehicles(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, vehicles)
}
