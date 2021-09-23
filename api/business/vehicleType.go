package business

import (
	"fmt"
	"github.com/Ricky-fight/car-admin-server/model/api/request"
	"github.com/Ricky-fight/car-admin-server/model/api/response"

	"github.com/Ricky-fight/car-admin-server/core"
	"github.com/Ricky-fight/car-admin-server/model/database"
	"github.com/Ricky-fight/car-admin-server/service/business"
	"github.com/Ricky-fight/car-admin-server/utilies"
	"github.com/gin-gonic/gin"
)

// VehicleTypeRentReq 车型基准租金
type VehicleTypeRentReq struct {
	Deposit     uint `json:"deposit" validate:"required"`     // 押金
	MonthlyRent uint `json:"monthlyRent" validate:"required"` // 租金
}

// VehicleTypeReq 车型
type VehicleTypeReq struct {
	Brand  string             `json:"brand"`                               // 厂牌
	TypeNo string             `json:"typeNo"`                              // 车型代号
	Series string             `json:"series" validate:"required"`          // 车系
	Rent   VehicleTypeRentReq `json:"vehicleTypeRent" validate:"required"` // FK 车型基准租金
}

// CreateVehicleType 创建车型
func CreateVehicleType(c *gin.Context) {
	// bind and validate data
	var r VehicleTypeReq
	if err := utilies.BindAndValid(&r, c); err != nil {
		return
	}
	// prepare data for service
	vt := database.VehicleType{
		Brand:  r.Brand,
		TypeNo: r.TypeNo,
		Series: r.Series,
		Rent: database.VehicleTypeRent{
			Deposit:     r.Rent.Deposit,
			MonthlyRent: r.Rent.MonthlyRent,
		},
	}
	fmt.Printf("vt: %v\n", vt)
	if err := business.CreateVehicleType(&vt); err != nil {
		core.FailWithErr(core.ERROR, err, c)
		return
	}
	core.Ok(c)
}

// DeleteVehicleTypeById 根据ID删除车型
func DeleteVehicleTypeById(c *gin.Context) {
	// get id
	id, err := utilies.GetIDFromParams(c)
	if err != nil {
		core.FailWithErr(core.ERROR, err, c)
	}
	// fmt.Printf("id in api: %v\n", id)
	if err := business.DeleteVehicleTypeById(id); err != nil {
		core.FailWithErr(core.ERROR, err, c)
		return
	}
	core.Ok(c)
}

// UpdateVehicleType 更新车型
func UpdateVehicleType(c *gin.Context) {
	var vt database.VehicleType
	if err := c.ShouldBindJSON(&vt); err != nil {
		core.FailWithBadRequest(c)
		return
	}
	if err := business.UpdateVehicleType(&vt); err != nil {
		core.FailWithErr(core.ERROR, err, c)
		return
	}
	core.Ok(c)
}

// GetVehicleTypeById 根据ID获取车型
func GetVehicleTypeById(c *gin.Context) {
	// get id
	id, err := utilies.GetIDFromParams(c)
	if err != nil {
		core.FailWithErr(core.ERROR, err, c)
		return
	}
	if vt, err := business.GetVehicleTypeById(id); err != nil {
		core.FailWithErr(core.ERROR, err, c)
		return
	} else {
		core.OkWithData(vt, c)
	}
}

func GetVehicleTypeList(c *gin.Context) {
	var info request.VehicleTypeSearch
	if err := c.ShouldBindJSON(&info); err != nil {
		core.FailWithBadRequest(c)
		return
	}
	if list, total, err := business.GetVehicleTypeList(info); err != nil {
		core.FailWithErr(core.ERROR, err, c)
		return
	} else {
		core.OkWithData(
			response.PageResult{
				List:     list,
				Total:    total,
				Page:     info.Page,
				PageSize: info.PageSize,
			},
			c,
		)
	}
}
