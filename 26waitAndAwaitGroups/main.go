package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race Condition")

	w8 := &sync.WaitGroup{}
	mut := &sync.RWMutex{}

	mut.RLock()
	var score = []int{0}
	mut.RUnlock()

	w8.Add(3)
	go func(w8 *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("One R")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		w8.Done()
	}(w8, mut)

	// w8.Add(1)
	go func(w8 *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Two R")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		w8.Done()
	}(w8, mut)

	// w8.Add(1)
	go func(w8 *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Three R")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		w8.Done()
	}(w8, mut)

	w8.Wait()

	fmt.Println(score)
}
