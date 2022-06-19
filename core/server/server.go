package server

import (
	"log"
	"net/http"
	"stashcurrency.dev/core/router"
)

func Run() {
	router := router.NewRouter()
	router.Intitialize()
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
