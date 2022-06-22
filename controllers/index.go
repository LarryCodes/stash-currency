package controllers

import (
	"log"
	"net/http"
	"stashcurrency.dev/core/currency"
	"stashcurrency.dev/core/initializer"
)

func Index(w http.ResponseWriter, r *http.Request) {

	currency_data, err := currency.LoadCurrencyList()
	if err != nil {
		log.Println("Failed to Loading currency_data failed!", err)
	}

	result_data := make(map[string]interface{})
	result_data["currency_data"] = currency_data

	err = initializer.Tpl.ExecuteTemplate(w, "index.gohtml", result_data)
	if err != nil {
		log.Println("Failed to execute template: ", err)
	}
}
