package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://thenutsup.com"

func main() {
	fmt.Println("The Nutsup web request")
	response, err := http.Get(url)
	checkNilErr(err)

	fmt.Printf("Response is of type: %T\n", response)

	defer response.Body.Close()

	dataBytes, err := io.ReadAll(response.Body)
	checkNilErr(err)
	fmt.Println("dataBytes", string(dataBytes))

}

func checkNilErr(err error) {
	if err != nil {
		panic((err))
	}
}
