package controllers

import (
	"log"
	"net/http"
	"stashcurrency.dev/core/currency"
	"stashcurrency.dev/core/initializer"
	"strconv"
)

func Convert(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	r.ParseForm()

	amount, _ := strconv.Atoi(r.FormValue("amount"))
	from := r.FormValue("from")
	to := r.FormValue("to")

	result_data := make(map[string]interface{})

	if amount != 0 && from != "" && to != "" {
		data, err := currency.Convert(amount, from, to)
		if err != nil {
			log.Println("Failed to convert currency!", err)
		}

		result_data["conversion_result"] = data
	}

	currency_data, err := currency.LoadCurrencyList()
	if err != nil {
		log.Println("Failed to Loading currency_data failed!", err)
	}

	result_data["currency_data"] = currency_data

	err = initializer.Tpl.ExecuteTemplate(w, "index.gohtml", result_data)
	if err != nil {
		log.Println("Failed to execute template: ", err)
	}

}
