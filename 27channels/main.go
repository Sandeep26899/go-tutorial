package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in golang")

	myCh := make(chan int, 2)
	w8 := &sync.WaitGroup{}

	// fmt.Println(<-myCh)
	// myCh <- 5

	w8.Add(2)
	//recieve only
	go func(ch <-chan int, w8 *sync.WaitGroup) {
		// close(myCh)
		val, isChOpen := <-myCh

		fmt.Println(isChOpen)
		fmt.Println(val)

		// fmt.Println(<-myCh)
		// fmt.Println(<-myCh)
		w8.Done()
	}(myCh, w8)

	//send only
	go func(ch chan<- int, w8 *sync.WaitGroup) {
		myCh <- 5
		close(myCh)
		w8.Done()
	}(myCh, w8)

	w8.Wait()

}
