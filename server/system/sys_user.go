package system

import (
	"server/global"
	"server/model/system"
)

type UserService struct{}

func (userService *UserService) Register(u system.SysUser) (err error, userInter system.SysUser) {
	err = global.DB.Create(&u).Error
	return err, u
}
