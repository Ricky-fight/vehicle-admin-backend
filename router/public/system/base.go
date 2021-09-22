package system

import (
	"github.com/Ricky-fight/car-admin-server/api/system"
	"github.com/gin-gonic/gin"
)

type BaseRouterGroup struct {
}

func (r BaseRouterGroup) Register(Router *gin.RouterGroup) {
	baseRouter := Router.Group("base")
	{
		baseRouter.POST("login", system.Login)
	}
}

var base = BaseRouterGroup{}
