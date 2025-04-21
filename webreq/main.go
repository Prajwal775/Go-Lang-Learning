package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to web request handling")
	// PerformGetRequest()
	PerformPostJsonRequest()
}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:3000/post"


	//fake json payload

	requestBody := strings.NewReader(`{"course": "golang", "price": 0, "platform": "udemy"}`)

	// http.Post(myurl, "application/json", requestBody)
	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {	
		panic(err)
	}
	content ,_ :=io.ReadAll(response.Body)
	fmt.Println("Response from server:", string(content))
	defer response.Body.Close()
	fmt.Println("Status Code:", response.StatusCode)

}

// func PerformGetRequest() {
// 	const myurl = "http://localhost:3000/get"
	

// 	response, err := http.Get(myurl)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer response.Body.Close()

// 	fmt.Println("Status Code:", response.StatusCode)
// 	fmt.Println("Content Length:", response.ContentLength)

// 	content, err := io.ReadAll(response.Body)
// 	if err != nil {
// 		panic(err)
// 	}

// 	//One way to print the response
// 	// fmt.Println("Response Body:", string(content))

// 	//Another way to print the response
// 	var responseString strings.Builder
// 	byteCount, _ := responseString.Write(content)
// 	fmt.Println("Number of bytes written:", byteCount)
// 	fmt.Println("Response Body:", responseString.String())
// }