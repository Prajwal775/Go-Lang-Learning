package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to web request handling")
	PerformGetRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:3000/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("Status Code:", response.StatusCode)
	fmt.Println("Content Length:", response.ContentLength)

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	//One way to print the response
	// fmt.Println("Response Body:", string(content))

	//Another way to print the response
	var responseString strings.Builder
	byteCount, _ := responseString.Write(content)
	fmt.Println("Number of bytes written:", byteCount)
	fmt.Println("Response Body:", responseString.String())
}