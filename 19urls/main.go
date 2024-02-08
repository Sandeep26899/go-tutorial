package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "http://localhost:3000/api/?name=sandeep&paymentid=san12345"

func main() {
	fmt.Println("Handling urls in golang")
	fmt.Println("myUrl", myUrl)

	// parsing
	result, _ := url.Parse(myUrl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Println(qparams["name"])

	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "http",
		Host:    "localhost:3000",
		Path:    "/api",
		RawPath: "name=sandeep",
	}

	anotherUrl := partsOfUrl.String()

	fmt.Println("anotherUrl: ", anotherUrl)
}
