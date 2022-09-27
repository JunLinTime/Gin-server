package initialize

import (
	"server/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GromMysqlByConfig(m config.DB) *gorm.DB {
	if m.Dbname == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(),
		DefaultStringSize:         191,
		SkipInitializeWithVersion: false,
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{}); err != nil {
		panic(err)
	} else {
		db.DB()
		return db
	}
}
