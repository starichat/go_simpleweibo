package database

import (
	"fmt"
	"go_simpleweibo/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lexkong/log"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

// 初始化数据库
func InitDB() *gorm.DB {
	db, err := gorm.Open(config.DBConfig.Connection, config.DBConfig.URL)
	if err != nil {
		log.Fatal("Database connection failed. Database url: "+config.DBConfig.URL+" error: ", err)
	} else {
		fmt.Println("\n\n-----------------------------GORM OPEN SUCCESS!------------------------------\n\n")
	}

	db = db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf-8;").AutoMigrate()

	db.LogMode(config.DBConfig.Debug)
	DB = db

	return DB
}
