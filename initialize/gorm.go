package initialize

import (
	"log"
	"os"
	"server/config"
	"server/model/system"

	"gorm.io/gorm"
)

func Gorm(m config.DB) *gorm.DB {
	return GromMysqlByConfig(m)
}

func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		system.SysUser{},
	)
	if err != nil {
		log.Println("register table failed", err)
		os.Exit(0)
	}
	log.Println("register table success")
}
