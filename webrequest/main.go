package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://www.wikipedia.org/"

func main() {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Printf("response is of type: %T\n", response)

	// dataBytes, err := io.ReadAll(response.Body)
	// if err != nil {
	// 	panic(err)
	// }

	// Limit to 1MB = 1,048,576 bytes
	limitedReader := io.LimitReader(response.Body, 1<<20) // 1MB

	dataBytes, err := io.ReadAll(limitedReader)
	if err != nil {
		panic(err)
	}


	content := string(dataBytes)
	fmt.Println(content)
}
