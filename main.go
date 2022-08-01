package main

import (
	"github.com/joho/godotenv"
	"log"
	"stashcurrency.dev/core/app"
)

func main() {

	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file: ", err)
	}

	// Run the app
	app.Run()
}
