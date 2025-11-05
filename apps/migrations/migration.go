package main

import (
	"fmt"
	"log"
	"mou-be/apps/config"
	"mou-be/apps/domain"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(".env")
	// Find and read the config file
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

}

func main() {

	db := config.DBConnect()

	// db.AutoMigrate(&domain.User{})
	// db.AutoMigrate(&domain.Session{})
	// db.AutoMigrate(&domain.Booth{})
	// db.AutoMigrate(&domain.Payment{})
	// db.AutoMigrate(&domain.PhotoResult{})
	// db.AutoMigrate(&domain.Voucher{})
	// db.AutoMigrate(&domain.Payment{})
	// db.AutoMigrate(&domain.Order{})
	// db.AutoMigrate(&domain.Voucher{})
	// db.AutoMigrate(&domain.VoucherTemplate{})
	// db.AutoMigrate(&domain.Booth{})
	db.AutoMigrate(&domain.Test{})

	fmt.Println("👍 Migration complete")
}
