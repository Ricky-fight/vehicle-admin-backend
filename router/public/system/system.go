package system

import (
	"github.com/Ricky-fight/car-admin-server/core"
	"github.com/gin-gonic/gin"
)

type SystemRouterGroup map[string]core.ApiRouter

var Path = "system"

func (r SystemRouterGroup) Register(Router *gin.RouterGroup) {
	systemRouter := Router.Group(Path)
	for _, v := range r {
		v.Register(systemRouter)
	}
}

var RouterGroup = SystemRouterGroup{
	base.Path: base,
}
