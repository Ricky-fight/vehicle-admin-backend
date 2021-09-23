package request

import "github.com/Ricky-fight/car-admin-server/model/database"

type VehicleTypeSearch struct {
	database.VehicleType
	PageInfo
}
