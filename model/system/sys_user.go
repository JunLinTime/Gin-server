package system

import "github.com/google/uuid"

type SysUser struct {
	UUID     uuid.UUID `json:"uuid" gorm:"comment:用户UUID"`
	Username string    `json:"userName" gorm:"comment:用户登录名"`
	Password string    `json:"-"  gorm:"comment:用户登录密码"`
}
