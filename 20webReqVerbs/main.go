package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb")
	// PerformGetRequest()
	// PerformPostRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:3000/users"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code:", response.StatusCode)
	fmt.Println("Content length:", response.ContentLength)

	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	fmt.Println(string(content))
	fmt.Println("byteCount is:", byteCount)
	fmt.Println(responseString.String())
}

func PerformPostRequest() {
	const myurl = "http://localhost:3000/users/check-username"

	body := strings.NewReader(`
		{
			"username": "sandeep"
		}
	`)

	response, err := http.Post(myurl, "application/json", body)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code:", response.StatusCode)
	fmt.Println("Content length:", response.ContentLength)
	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	fmt.Println(string(content))
	fmt.Println("byteCount is:", byteCount)
	fmt.Println(responseString.String())
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:3000/users/check-username"

	data := url.Values{}
	data.Add("username", "sandeep")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code:", response.StatusCode)
	fmt.Println("Content length:", response.ContentLength)
	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	fmt.Println(string(content))
	fmt.Println("byteCount is:", byteCount)
	fmt.Println(responseString.String())
}
