package database

import (
	"log"

	"github.com/spf13/viper"
)

func Initdb() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error membaca file .env: %v", err)
	}
}
