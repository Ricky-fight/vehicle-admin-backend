package business

import (
	"fmt"
	"github.com/Ricky-fight/car-admin-server/global"
	"github.com/Ricky-fight/car-admin-server/model/api/request"
	"github.com/Ricky-fight/car-admin-server/model/database"
	"github.com/kr/pretty"
	"gorm.io/gorm"
)

// CreateVehicleType 创建车型
func CreateVehicleType(vt *database.VehicleType) error {
	// var rst database.VehicleType
	err := global.DB.Create(vt).Error
	return err
}

// DeleteVehicleTypeById 删除车型
func DeleteVehicleTypeById(id uint) error {
	var vt database.VehicleType
	vt.ID = id
	fmt.Printf("vt.ID: %v\n", vt.ID)
	// TODO 增加判断 车型下是否还有车辆，有的话不能删除，只能修改
	// 查询车型和租金
	db := global.DB
	db.Preload("Rent").Find(&vt)
	fmt.Printf("vt: %v\n", vt)
	{
		// 删除车型
		db.Delete(&vt)
		// 删除车型租金记录
		db.Delete(&vt.Rent)
	}
	return nil
}

// UpdateVehicleType 更新车型
func UpdateVehicleType(vt *database.VehicleType) error {
	if vt.ID == 0 {
		return fmt.Errorf("ID hasn't given")
	}
	pretty.Println(vt.Rent)
	if err := global.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&vt).Error; err != nil {
		return err
	}
	return nil
}

// GetVehicleTypeById 根据ID查询车型
func GetVehicleTypeById(id uint) (vt database.VehicleType, err error) {
	err = global.DB.Preload("Rent").First(&vt, id).Error
	return
}

// GetVehicleTypeList 查询车型列表
func GetVehicleTypeList(info request.VehicleTypeSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.DB.Model(&database.VehicleType{}).Preload("Rent")
	var vehicleTypes []database.VehicleType
	// 如果有条件搜索
	//下方会自动创建搜索语句

	if info.TypeNo != "" {
		db = db.Where("`license_no` LIKE ?", "%"+info.TypeNo+"%")
	}
	if info.Series != "" {
		db = db.Where("`vin` LIKE ?", "%"+info.Series+"%")
	}
	if info.Brand != "" {
		db = db.Where("`engine_no` LIKE ?", "%"+info.Brand+"%")
	}
	err = db.Count(&total).Limit(limit).Offset(offset).Find(&vehicleTypes).Error
	return vehicleTypes, total, err
}
