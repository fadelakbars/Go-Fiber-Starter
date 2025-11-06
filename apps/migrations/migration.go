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

	db.AutoMigrate(&domain.User{})

	fmt.Println("👍 Migration complete")
}
