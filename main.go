package main

import (
	"github.com/joho/godotenv"
	"log"
	"stashcurrency.dev/core/app"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file: ", err)
	}

	app.Run()
}
