package business

import (
	"github.com/Ricky-fight/car-admin-server/api/business"
	"github.com/gin-gonic/gin"
)

type carTypeRouterGroup struct {
	Path string
}

func (r carTypeRouterGroup) Register(Router *gin.RouterGroup) {
	carTypeRouter := Router.Group("")
	{
		path := carType.Path
		carTypeRouter.POST(path, business.CreateVehicleType)
		carTypeRouter.DELETE(path+"/:id", business.DeleteVehicleTypeById)
		carTypeRouter.PUT(path, business.UpdateVehicleType)
		carTypeRouter.GET(path+"/:id", business.GetVehicleTypeById)
		carTypeRouter.GET(path, business.GetVehicleTypeList)
	}
}

var carType = carTypeRouterGroup{
	"cartypes",
}
