package initialize

import "github.com/Ricky-fight/car-admin-server/global"

func Init() {
	// 加载配置
	Viper()
	// 连接数据库并挂载到全局变量
	global.DB = DB()
	// 初始化路由
	global.ROUTER = Router()
}
