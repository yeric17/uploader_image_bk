package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	PROD_HOST, MODE, DEV_HOST, API_PORT, FT_DEV_HOST, FT_PROD_HOST, HOST, FT_HOST string
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}
	PROD_HOST = os.Getenv("PROD_HOST")
	MODE = os.Getenv("MODE")
	DEV_HOST = os.Getenv("DEV_HOST")
	API_PORT = os.Getenv("API_PORT")
	os.Setenv("PORT", API_PORT)
	FT_PROD_HOST = os.Getenv("FT_PROD_HOST")
	FT_DEV_HOST = os.Getenv("FT_DEV_HOST")
	if MODE == "dev" {
		HOST = fmt.Sprintf("%s:%s", DEV_HOST, API_PORT)
		FT_HOST = FT_DEV_HOST
	} else {
		HOST = PROD_HOST
		FT_HOST = FT_PROD_HOST
	}
}
