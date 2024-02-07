package main

import "fmt"

func main() {
	fmt.Println("loops in golang")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	fmt.Println("Days", days)

	//First way
	for d := 0; d < len(days); d++ {
		fmt.Println("day", days[d])
	}

	fmt.Println("-----------Second-----------")

	//Second way (not recommended)
	for i := range days {
		fmt.Println("day", days[i])
	}

	fmt.Println("-----------Third-----------")

	// Third way
	for _, day := range days {
		fmt.Printf("values is %v\n", day)
	}

	fmt.Println("-----------Fourth-----------")

	// Fourth way
	rougueValue := 1

	for rougueValue <= 10 {

		if rougueValue == 2 {
			goto lco
		}

		if rougueValue == 5 {
			rougueValue++
			continue
		}
		fmt.Println("rougueValue", rougueValue)
		rougueValue++
	}

lco:
	fmt.Println("Jumping here")
}
