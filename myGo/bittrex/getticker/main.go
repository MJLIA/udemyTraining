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

type GetTicker struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Result  struct {
		Bid  float64 `json:"Bid"`
		Ask  float64 `json:"Ask"`
		Last float64 `json:"Last"`
	} `json:"result"`
}

func main() {

	response, err := http.Get("https://bittrex.com/api/v1.1/public/getticker")

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

	var responseObject GetTicker
	json.Unmarshal(responseData, &responseObject)
	fmt.Println(responseObject.Result)

	//for i := 0; i < len(responseObject.Result); i++ {
	//	fmt.Println(responseObject.Result[i].Currency, responseObject.Result[i].CurrencyLong)

}
