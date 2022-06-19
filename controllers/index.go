package controllers

import (
	"log"
	"net/http"
	"stashcurrency.dev/core/initializer"
)

func Index(w http.ResponseWriter, r *http.Request) {
	err := initializer.Tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		log.Println("Failed to execute template: ", err)
	}
}

func loadCurrencyList() {
	
}