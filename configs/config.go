package configs

import (
	"fmt"
	"log"
	"nahwudasar-be/helper"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DBConn     *gorm.DB
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
)

func ConnectDB() {
	// fmt.Println(os.Getenv("DB_HOST"))
	DbHost = helper.GetEnv("DB_HOST")
	DbPort = helper.GetEnv("DB_PORT")
	DbUser = helper.GetEnv("DB_USER")
	DbPassword = helper.GetEnv("DB_PASSWORD")
	DbName = helper.GetEnv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=UTC", DbUser, DbPassword, DbHost, DbPort, DbName)
	// dsn := fmt.Sprintf("%s:%s@%s:%s/%s?charset=utf8mb4&parseTime=True&loc=UTC", DbUser, DbPassword, DbHost, DbPort, DbName)

	my := mysql.New(mysql.Config{DSN: dsn})

	db, err := gorm.Open(my, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		// log.Printf("Failed to connect to database. %v\n", err)
		log.Printf(dsn, err)
		os.Exit(2)
	}

	log.Println("Connected")
	DBConn = db

}
