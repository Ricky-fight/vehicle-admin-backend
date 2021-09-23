package system

import (
	"github.com/Ricky-fight/car-admin-server/api/system"
	"github.com/gin-gonic/gin"
)

type BaseRouterGroup struct {
	Path string
}

func (r BaseRouterGroup) Register(Router *gin.RouterGroup) {
	baseRouter := Router.Group(base.Path)
	{
		baseRouter.POST("login", system.Login)
		baseRouter.POST("register", system.Register)
	}
}

var base = BaseRouterGroup{
	Path: "base",
}
