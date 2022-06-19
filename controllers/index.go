package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"stashcurrency.dev/core/initializer"
)

type currencies struct {
	Currencies map[string]string `json:"currencies"`
}

func Index(w http.ResponseWriter, r *http.Request) {

	currency_data, err := loadCurrencyList()
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

func loadCurrencyList() (*currencies, error) {
	url := "https://api.apilayer.com/currency_data/list"

	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("apikey", "0en0LKbSXjcIPe2O2k5KiudiCwUPE5EN")

	if err != nil {
		return nil, err
	}

	result, err := client.Do(request)
	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return nil, err
	}

	var currency_data currencies
	if err = json.Unmarshal(body, &currency_data); err != nil {
		return nil, err
	}

	return &currency_data, nil
}
