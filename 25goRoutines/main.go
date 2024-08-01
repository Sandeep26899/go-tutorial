package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

var w8 sync.WaitGroup //pointers
var mut sync.Mutex    //pointers

func main() {
	// go greeter("Hello")
	// greeter("World")

	websitelist := []string{
		"https://thenutsup.com",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _, web := range websitelist {
		go getStatusCode(web)
		w8.Add(1)
	}

	w8.Wait()
	fmt.Println(signals)
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer w8.Done()

	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OOPS in endpoint:", endpoint)
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()

		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}

}
