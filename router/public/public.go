package public

import (
	"github.com/Ricky-fight/car-admin-server/core"
	"github.com/Ricky-fight/car-admin-server/router/public/business"
	"github.com/Ricky-fight/car-admin-server/router/public/system"
	"github.com/gin-gonic/gin"
)

type PublicRouterGroup map[string]core.ApiRouter

func (r PublicRouterGroup) Register(Router *gin.RouterGroup) {
	publicRouter := Router.Group("")
	for _, v := range r {
		v.Register(publicRouter)
	}
}

var RouterGroup = PublicRouterGroup{
	system.Path:   system.RouterGroup,
	business.Path: business.RouterGroup,
}
