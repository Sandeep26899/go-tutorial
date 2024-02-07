package main

import "fmt"

func main() {
	fmt.Println("Functions in golang")
	greeterTwo()
	greeter()
	result := adder(5, 5)

	fmt.Println("Result is", result)

	resultProAdder, message := proAdder(7, 5, 5, 7)

	fmt.Println("ProAdder result is", resultProAdder)
	fmt.Println("ProAdder message is", message)
}

func greeterTwo() {
	fmt.Println("Another Func")
}

func greeter() {
	fmt.Println("Hello")
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "This is result"
}
