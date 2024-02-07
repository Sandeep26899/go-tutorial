package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or parent

	user := User{"Sandeep", "sandeepsj26899@gmail.com", true, 24}

	fmt.Println(user)

	fmt.Printf("sandeep details are: %+v\n", user)
	fmt.Printf("Name is %v and email is %v\n", user.Name, user.Email)
	user.GetStatus()
	user.NewEmail()
	fmt.Printf("Name is %v and email is %v\n", user.Name, user.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active", u.Status)
}

func (u User) NewEmail() {
	u.Email = "test@go.dev"
	fmt.Println("New Email", u.Email)
}
