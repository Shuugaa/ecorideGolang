package main

import (
	"ecoride/database"
	"ecoride/router"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	r := router.ServeRouter()
	database.CreateTableUsers()
	database.CreateTableSessions()
	r.Run(":3000")
}
