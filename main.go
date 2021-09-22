package main

import (
	"github.com/Ricky-fight/car-admin-server/core"
	"github.com/Ricky-fight/car-admin-server/initialize"
)

func main() {
	// 初始化
	initialize.Init()
	// 启动服务器
	core.RunServer()
}
