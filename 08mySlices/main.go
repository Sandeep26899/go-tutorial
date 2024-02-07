package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices module")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitList is %T \n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3], "")

	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 955
	highScores[2] = 465
	highScores[3] = 867
	// highScores[4] = 514

	highScores = append(highScores, 555, 796, 888)

	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))

	sort.Ints(highScores)
	fmt.Println(highScores)

	var courses = []string{"reactjs", "javascript", "swift", "php", "python"}
	fmt.Println("courses:", courses)
	var index int = 0
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
