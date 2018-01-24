package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Networks struct {
	address string `xml:"address"`
	Host    []Host `xml:"address"`
}

func main() {

	xmlFile, err := os.Open("Usersnet.xml")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully loaded Usersnet.xml")
	defer xmlFile.Close()

	responseData, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		log.Fatal(err)

		var responseObject Networks
		xml.Unmarshal(responseData, &responseObject)
		fmt.Println(responseObject.Result)

		for i := 0; i < len(responseObject.Result); i++ {
			fmt.Println(responseObject.Result[i].host, responseObject.Result[i].address)

		}
	}
}
