package system

import (
	"fmt"

	"github.com/Ricky-fight/car-admin-server/global"
	"github.com/Ricky-fight/car-admin-server/model/database"
)

func Login(u *database.User) error {
	rst := &database.User{}
	if global.DB.Where(u).First(rst); rst.ID == 0 {
		return fmt.Errorf("login service failed")
	}
	return nil
}
func Register(u *database.User) error {
	var ret database.User
	ret.Account = u.Account
	if err := global.DB.Limit(1).Find(&ret).Error; err != nil {
		// 遇到未知错误
		return fmt.Errorf("create account failed, unknown error when finding exist account record")
	} else {
		if ret.Password != "" {
			// 找到已注册记录
			return fmt.Errorf("create account failed, please check if the account '%v' already exists", u.Account)
		} else {
			// 未找到记录，可以注册
			if err := global.DB.Create(u).Error; err != nil {
				// 遇到未知错误
				return fmt.Errorf("create account failed, unknown error")
			}
		}
	}
	return nil
}
