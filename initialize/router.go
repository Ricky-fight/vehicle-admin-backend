package initialize

import (
	"github.com/Ricky-fight/car-admin-server/router/public"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	engine := gin.Default()
	router := engine.Group("")
	public.RouterGroup.Register(router)
	return engine
}
