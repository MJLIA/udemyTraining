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

type GetMarketSum struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Result  []struct {
		MarketName        string      `json:"MarketName"`
		High              float64     `json:"High"`
		Low               float64     `json:"Low"`
		Volume            float64     `json:"Volume"`
		Last              float64     `json:"Last"`
		BaseVolume        float64     `json:"BaseVolume"`
		TimeStamp         string      `json:"TimeStamp"`
		Bid               float64     `json:"Bid"`
		Ask               float64     `json:"Ask"`
		OpenBuyOrders     int         `json:"OpenBuyOrders"`
		OpenSellOrders    int         `json:"OpenSellOrders"`
		PrevDay           float64     `json:"PrevDay"`
		Created           string      `json:"Created"`
		DisplayMarketName interface{} `json:"DisplayMarketName"`
	} `json:"result"`
}

func main() {

	response, err := http.Get("https://bittrex.com/api/v1.1/public/getmarketsummaries")

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

	var responseObject GetMarketSum
	json.Unmarshal(responseData, &responseObject)
	fmt.Println(responseObject.Result)

	for i := 0; i < len(responseObject.Result); i++ {
		fmt.Println(responseObject.Result[i].MarketName, responseObject.Result[i].OpenBuyOrders, responseObject.Result[i].BaseVolume)

	}
}
