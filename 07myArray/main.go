package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golangs")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[2] = "Mango"

	fmt.Println("Fruit List:", fruitList)
	fmt.Println("Fruit List Length:", len(fruitList))

	var vegList = [5]string{"Potato", "Beans", "Mushroom"}

	fmt.Println("Vegetables List:", vegList)
	fmt.Println("Vegetables List Length:", len(vegList))
}
