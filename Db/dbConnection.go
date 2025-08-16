package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var GormDB *gorm.DB

func Db_Connection() error {
	// dsn := "host=localhost user=postgres password=guru dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	var err error

	dsn := "root:root@tcp(localhost:3306)/dev?charset=utf8mb4&parseTime=True&loc=Local"
	GormDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	sqlDB, err := GormDB.DB()
	if err != nil {
		fmt.Println("error getting generic DB from GORM:", err)
		return err
	}

	if err := sqlDB.Ping(); err != nil {
		fmt.Println("database ping failed:", err)
		return err
	}

	fmt.Println("connected successfully")
	return nil
}
