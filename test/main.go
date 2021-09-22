package main

import (
	"fmt"

	"github.com/Ricky-fight/car-admin-server/global"
	"github.com/Ricky-fight/car-admin-server/initialize"
	"github.com/Ricky-fight/car-admin-server/model/database"
)

func main() {
	initialize.Init()
	u := database.User{
		Account:  "111",
		Password: "11",
	}
	var u1 database.User
	err := global.DB.FirstOrCreate(&u1, u).Error
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("u1: %v\n", u1)
}
