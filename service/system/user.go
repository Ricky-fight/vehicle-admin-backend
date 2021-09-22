package system

import (
	"github.com/Ricky-fight/car-admin-server/global"
	"github.com/Ricky-fight/car-admin-server/model/database"
)

func Login(u *database.User) bool {
	rst := &database.User{}
	if global.DB.Where(u).First(rst); rst.ID != 0 {
		return true
	}
	return false

}
