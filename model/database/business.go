package database

import (
	"time"
)

type Vehicle struct {
	// CID string
	LicenseNo string // 车牌号
	VIN       string // 车架号/车辆识别代号
	EngineNo  string // 发动机号
	// UseNature     uint
	VehicleTypeID         uint                // FK 车型
	VehicleType           VehicleType         // 车型实体
	EnrollDate            time.Time           // 初登日期
	VehicleRentContractID uint                // FK 车辆租赁合同
	VehicleRentContract   VehicleRentContract // FK 车辆租赁合同
	Model
}

type VehicleType struct {
	Brand                 string          // 厂牌
	TypeNo                string          // 车型代号
	Series                string          // 车系
	VehicleRentContractID uint            // FK 车型基准租金
	VehicleTypeRent       VehicleTypeRent // FK 车型基准租金
	Vehicles              []Vehicle       // FK 下属车辆
	Model
}
type VehicleTypeRent struct {
	VehicleTypeID uint // FK 车型
	Deposit       uint // 押金
	MonthlyRent   uint // 租金
}
type Client struct {
	// Type           bool   // 客户类型,先默认全部为司机客户，后续添加公司客户管理再增加此子段
	Name string // 客户名称
	// CredentialType int
	CredentialNo        string                // 证件号码，后续为区别证件类型会增加字段
	ClientContact       ClientContact         // FK 联系方式
	VehicleRentContract []VehicleRentContract // FK hasmany 车辆租赁合同
	Model
}

type ClientContact struct {
	ClientID    uint   // FK 客户
	Phone       string // 电话
	PhoneBackup string // 备用电话
	Address     string // 联系地址
	Model
}

type VehicleRentContract struct {
	ClientID         uint      // FK 客户
	VehicleID        uint      // FK 承租车辆
	SignedDate       time.Time // 合同签订日期
	EffectiveDate    time.Time // 合同起效日期
	TerminatedDate   time.Time // 合同终止日期
	Deposit          uint      // 押金
	MonthlyRent      uint      // 月租金
	BillingCycle     bool      // false月付/true周付
	FirstBillingDate time.Time // 首次付租日期
}
