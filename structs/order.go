package structs

import (
	"time"
)

type Order struct {
	ID             int       `json:"id"`
	OrderEntryDate time.Time `json:"orderEntryDate"`
	VehicleId      int       `json:"vehicleId"`
	CustomerName   string    `json:"customerName"`
}
