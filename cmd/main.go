package main

import (
	"github.com/joho/godotenv"
	"log"
	"realTimeCryptoPrice/api/repository"
)

func init() {

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
}

func main() {

	result, err := repository.CoinMarketCapAPI("cardano")
	if err != nil {
		panic(err)
	}
}
