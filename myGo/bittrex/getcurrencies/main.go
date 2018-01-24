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

type Currencies struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Result  []struct {
		Currency        string      `json:"Currency"`
		CurrencyLong    string      `json:"CurrencyLong"`
		MinConfirmation int         `json:"MinConfirmation"`
		TxFee           float64     `json:"TxFee"`
		IsActive        bool        `json:"IsActive"`
		CoinType        string      `json:"CoinType"`
		BaseAddress     interface{} `json:"BaseAddress"`
	} `json:"result"`
}

func main() {

	response, err := http.Get("https://bittrex.com/api/v1.1/public/getcurrencies")

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

	var responseObject Currencies
	json.Unmarshal(responseData, &responseObject)
	fmt.Println(responseObject.Result)

	for i := 0; i < len(responseObject.Result); i++ {
		fmt.Println(responseObject.Result[i].Currency, responseObject.Result[i].CurrencyLong)

	}
}
