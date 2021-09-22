package initialize

import (
	"fmt"

	"github.com/Ricky-fight/car-admin-server/global"
	"github.com/Ricky-fight/car-admin-server/model/database"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DB() (db *gorm.DB) {
	m := global.CONFIG.Mysql
	dsn := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		m.User,
		m.Password,
		m.Host,
		m.Port,
		m.Dbname,
	)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// 自动迁移数据表
	MigrateTables(db)
	return
}
func MigrateTables(db *gorm.DB) {
	db.AutoMigrate(
		// system
		database.User{},

		// business
		database.Vehicle{},
		database.VehicleType{},
		database.VehicleTypeRent{},
		database.Client{},
		database.ClientContact{},
		database.VehicleRentContract{},
	)
}
