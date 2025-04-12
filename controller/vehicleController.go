package controller

import (
	"fmt"
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

func GetVehicles(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, vehicles)
}

func PostVehicle(c *gin.Context) {
	var newVehicle structs.Vehicle

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newVehicle); err != nil {
		fmt.Print(err)
		return
	}

	// Add the new album to the slice.
	vehicles = append(vehicles, newVehicle)
	c.IndentedJSON(http.StatusCreated, newVehicle)
}
