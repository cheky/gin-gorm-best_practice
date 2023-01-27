package config

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var mysql_DB *gorm.DB

func InitMysql() *gorm.DB {
	mysql_DB = ConnectMysql()
	return mysql_DB
}
func ConnectMysql() *gorm.DB {
	mode := os.Getenv("MODE")
	MYSQL_USERNAME := os.Getenv("LIVE-MYSQL-USERNAME")
	MYSQL_PASSWORD := os.Getenv("LIVE-MYSQL-PASSWORD")
	MYSQL_DATABASE := os.Getenv("LIVE-MYSQL-DATABASE")
	MYSQL_HOST := os.Getenv("LIVE-MYSQL-HOST")
	MYSQL_PORT := os.Getenv("LIVE-MYSQL-PORT")
	if mode == "sandbox" {
		MYSQL_USERNAME = os.Getenv("SANDBOX-MYSQL-USERNAME")
		MYSQL_PASSWORD = os.Getenv("SANDBOX-MYSQL-PASSWORD")
		MYSQL_DATABASE = os.Getenv("SANDBOX-MYSQL-DATABASE")
		MYSQL_HOST = os.Getenv("SANDBOX-MYSQL-HOST")
		MYSQL_PORT = os.Getenv("SANDBOX-MYSQL-PORT")
	}
	var err error
	dsn := MYSQL_USERNAME + ":" + MYSQL_PASSWORD + "@tcp" + "(" + MYSQL_HOST + ":" + MYSQL_PORT + ")/" + MYSQL_DATABASE + "?" + "parseTime=true&loc=Local"
	fmt.Println("(MYSQL) dsn : ", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println("(MYSQL) Error connecting to database : error=%v", err)
		return nil
	}

	return db
}
