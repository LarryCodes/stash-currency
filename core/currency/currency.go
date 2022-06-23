package currency

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type CurrencyList struct {
	Currencies map[string]string `json:"currencies"`
}

type ConversionResult struct {
	Result     float64         `json:"result"`
	Success    bool            `json:"success"`
	Date       string          `json:"date"`
	Historical bool            `json:"historical"`
	Info       ConversionInfo  `json:"info"`
	Query      ConversionQuery `json:"query"`
}

type ConversionInfo struct {
	Quote     float64 `json:"quote"`
	Timestamp int64   `json:"timestamp"`
}

type ConversionQuery struct {
	Amount float64 `json:"amount"`
	From   string  `json:"from"`
	To     string  `json:"to"`
}

func LoadCurrencyList() (*CurrencyList, error) {
	url := "https://api.apilayer.com/currency_data/list"

	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("apikey", os.Getenv("API_KEY"))

	if err != nil {
		return nil, err
	}

	result, err := client.Do(request)
	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return nil, err
	}

	var currency_data CurrencyList
	if err = json.Unmarshal(body, &currency_data); err != nil {
		return nil, err
	}

	return &currency_data, nil
}

func Convert(amount int, from string, to string) (*ConversionResult, error) {

	url := fmt.Sprintf("https://api.apilayer.com/currency_data/convert?to=%s&from=%s&amount=%d", to, from, amount)
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("apikey", os.Getenv("API_KEY"))

	if err != nil {
		return nil, err
	}

	result, err := client.Do(req)
	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(body))
	var data ConversionResult
	if err = json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return &data, nil

}
