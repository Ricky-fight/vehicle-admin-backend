package business

import (
	"fmt"
	"github.com/Ricky-fight/car-admin-server/global"
	"github.com/Ricky-fight/car-admin-server/model/database"
	"gorm.io/gorm"
)

// CreateVehicle 创建车辆
func CreateVehicle(vehicle *database.Vehicle) error {
	err := global.DB.Create(vehicle).Error
	return err
}

// DeleteVehicleById 删除车辆,不推荐用
func DeleteVehicleById(id uint) error {
	if id == 0 {
		return fmt.Errorf("ID hasn't given")
	}

	var vehicle database.Vehicle
	db := global.DB.Model(&vehicle).Preload("Contracts")
	// 删除车辆
	db.Find(&vehicle, id)
	if len(vehicle.Contracts) != 0 {
		return fmt.Errorf("该车辆记录下有合同，请先删除关联合同")
	}
	db.Delete(&vehicle, id)
	// db.Delete(&vehicle.Contracts)
	return nil
}

// UpdateVehicle 更新车辆
func UpdateVehicle(vehicle *database.Vehicle) error {
	if vehicle.ID == 0 {
		return fmt.Errorf("ID hasn't given")
	}
	if err := global.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&vehicle).Error; err != nil {
		return err
	}
	return nil
}

// GetVehicleById 根据ID查询车辆
func GetVehicleById(id uint) (vt database.Vehicle, err error) {
	err = global.DB.Preload("Rent").First(&vt, id).Error
	return
}

// GetVehicleList 查询车辆列表
//func GetVehicleList(info request.VehicleSearch) (list interface{}, total int64, err error) {
//	limit := info.PageSize
//	offset := info.PageSize * (info.Page - 1)
//	// 创建db
//	db := global.DB.Model(&database.Vehicle{}).Preload("Rent")
//	var vehicles []database.Vehicle
//	// 如果有条件搜索
//	//下方会自动创建搜索语句
//
//	if info.TypeNo != "" {
//		db = db.Where("`license_no` LIKE ?", "%"+info.TypeNo+"%")
//	}
//	if info.Series != "" {
//		db = db.Where("`vin` LIKE ?", "%"+info.Series+"%")
//	}
//	if info.Brand != "" {
//		db = db.Where("`engine_no` LIKE ?", "%"+info.Brand+"%")
//	}
//	err = db.Count(&total).Limit(limit).Offset(offset).Find(&vehicles).Error
//	return vehicles, total, err
//}
