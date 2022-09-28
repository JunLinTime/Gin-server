package main

import (
	"log"
	"server/config"
	"server/global"
	"server/initialize"
)

func main() {
	dbConfig := config.DB{
		Username: "root",
		Password: "test",
		Path:     "127.0.0.1",
		Port:     "3306",
		Dbname:   "gin",
		Config:   "charset=utf8mb4&parseTime=True&loc=Local",
	}
	global.DB = initialize.GromMysqlByConfig(dbConfig)
	if global.DB != nil {
		initialize.RegisterTables(global.DB)
	} else {
		log.Println("global DB is nil")
	}
	router := initialize.Routers()
	router.Run(":9124")
}
