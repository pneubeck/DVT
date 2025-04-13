package VehicleController

import (
	"fmt"
	"net/http"
	"strconv"

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

func GetVehicleById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range vehicles {
		if strconv.Itoa(a.ID) == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "vehicle not found"})
}

func PostVehicle(c *gin.Context) {
	var newVehicle structs.Vehicle

	if err := c.BindJSON(&newVehicle); err != nil {
		fmt.Print(err)
		return
	}

	vehicles = append(vehicles, newVehicle)
	c.IndentedJSON(http.StatusCreated, newVehicle)
}
