package config

import (
	"fmt"
	"strconv"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func DBConnect() *gorm.DB {
	host := viper.Get("DB_HOST")
	port, _ := strconv.Atoi(viper.Get("DB_PORT").(string))
	user := viper.Get("DB_USER")
	password := "8ybMXUcgcPF8YT$m2g90" //viper.Get("DB_PASSWORD")
	dbname := viper.Get("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Database connected ... !")
	return db
}
