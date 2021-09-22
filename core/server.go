package core

import (
	"github.com/Ricky-fight/car-admin-server/global"
)

func RunServer() {
	port := global.CONFIG.System.Port
	router := global.ROUTER
	router.Run(":" + port)
}
