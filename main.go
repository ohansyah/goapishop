package main

import (
	"api_olshop/database"
	"api_olshop/models"
	"api_olshop/routes"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func main() {

	// Load ENV
	GetEnvType := GetEnvType("DEVELOPMENT_TYPE")
	fmt.Println("DEVELOPMENT_TYPE : ", GetEnvType)

	viper.SetConfigFile("./configs/" + GetEnvType + ".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	// Database connection
	db := database.ConnectToDB()

	// migration
	db.AutoMigrate(&models.Contact{})
	defer db.Close()

	// routing
	routes.HandleRequest()
}

// GetEnvType get type project. local develop production
func GetEnvType(key string) string {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	return value
}
