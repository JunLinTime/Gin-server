package main

import (
	"server/config"
	"server/global"
	"server/initialize"
)

func main() {
	dbConfig := config.DB{
		Username: "root",
		Password: "test",
		Path:     "127.0.0.1",
		Port:     "9123",
		Config:   "charset=utf8mb4&parseTime=True&loc=Local",
	}
	global.DB = initialize.GromMysqlByConfig(dbConfig)
	router := initialize.Routers()
	router.Run(":9124")
	if global.DB != nil {
		initialize.RegisterTables(global.DB)
		db, _ := global.DB.DB()
		defer db.Close()
	}
}
