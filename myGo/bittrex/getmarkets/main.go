package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Bittrex getmarkets struct

type Markets struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Result  []struct {
		MarketCurrency     string  `json:"MarketCurrency"`
		BaseCurrency       string  `json:"BaseCurrency"`
		MarketCurrencyLong string  `json:"MarketCurrencyLong"`
		BaseCurrencyLong   string  `json:"BaseCurrencyLong"`
		MinTradeSize       float64 `json:"MinTradeSize"`
		MarketName         string  `json:"MarketName"`
		IsActive           bool    `json:"IsActive"`
		Created            string  `json:"Created"`
	} `json:"result"`
}

func main() {

	response, err := http.Get("https://bittrex.com/api/v1.1/public/getmarkets")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)

	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Print out entire response Body
	//	fmt.Println(string(responseData))

	var responseObject Markets
	json.Unmarshal(responseData, &responseObject)
	fmt.Println(responseObject.Result)

	for i := 0; i < len(responseObject.Result); i++ {
		fmt.Println(responseObject.Result[i].MarketCurrency, responseObject.Result[i].BaseCurrency, responseObject.Result[i].MarketCurrencyLong)

	}
}
