package initialize

import (
	"log"
	"server/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GromMysqlByConfig(m config.DB) *gorm.DB {
	if m.Dbname == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN: m.Dsn(),
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{}); err != nil {
		log.Println(err)
		panic(err)
	} else {
		db.DB()
		return db
	}
}
