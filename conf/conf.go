package conf

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Init() {
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "development"
	}

	err := godotenv.Load(".env." + env)
	if err != nil {
		log.Fatal(err)
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}
