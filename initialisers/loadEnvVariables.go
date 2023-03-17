package initialisers

import (
	"log"

	"github.com/joho/godotenv"
)

// every variable that's inside the .env file is like an input to a flag when executing main.go
func LoadEnvVariables() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file.")
	}
}
