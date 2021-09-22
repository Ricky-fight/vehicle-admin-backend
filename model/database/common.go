package database

import (
	"gorm.io/gorm"
)

// 设计规范：禁用gorm.Model,所有表里请自建主键，用带有公司标识或具有规律的主键（完成项目后再实施，现在先用自带主键）

// Model a basic GoLang struct which includes the following fields: ID, CreatedAt, UpdatedAt, DeletedAt
// 基础 Model 不带ID，需要额外建立一个主键 e.g. `gorm:"primarykey"`
// type BaseModel struct {
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// 	DeletedAt gorm.DeletedAt `gorm:"index"`
// }

type Model gorm.Model
