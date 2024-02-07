package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or parent

	user := User{"Sandeep", "sandeepsj26899@gmail.com", true, 24}

	fmt.Println(user)

	fmt.Printf("sandeep details are: %+v\n", user)
	fmt.Printf("Name is %v and email is %v\n", user.Name, user.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
