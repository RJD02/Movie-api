package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetConfig() map[string]string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	config := make(map[string]string)

	if val, ok := os.LookupEnv("MONGODB_USER"); ok {
		config["MONGODB_USER"] = val
	}

	if val, ok := os.LookupEnv("MONGODB_PASSWORD"); ok {
		config["MONGODB_PASSWORD"] = val
	}

	return config
}
