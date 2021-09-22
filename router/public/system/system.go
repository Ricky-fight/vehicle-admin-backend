package system

import (
	"github.com/Ricky-fight/car-admin-server/core"
	"github.com/gin-gonic/gin"
)

type SystemRouterGroup map[string]core.ApiRouter

func (r SystemRouterGroup) Register(Router *gin.RouterGroup) {
	systemRouter := Router.Group("system")
	for _, v := range r {
		v.Register(systemRouter)
	}
}

var RouterGroup = SystemRouterGroup{
	"base": base,
}
