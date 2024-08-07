package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	// "math/rand"
)

func main() {
	fmt.Println("Maths in golang")

	var numberOne int = 2
	var numberTwo float64 = 4

	fmt.Println("The sum is:", numberOne+int(numberTwo))

	//random number with math
	// fmt.Println(rand.Intn(6))

	// random with crypto
	randNum, err := rand.Int(rand.Reader, big.NewInt(5))

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("randNum", randNum)
}
