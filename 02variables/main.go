package main

import "fmt"

const LoginToken string = "jbsdhvbsgv" // pubic keyword

func main() {
	var name string = "Sandeep"
	fmt.Println(name)
	fmt.Printf("Variable name type: %T \n", name)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable isLoggedIn type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable smallVal type: %T \n", smallVal)

	var smallFloat float32 = 255.4555487451514
	fmt.Println(smallFloat)
	fmt.Printf("Variable smallFloat type: %T \n", smallFloat)

	//for default value
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable anotherVariable type: %T \n", anotherVariable)

	//implicit type

	var website = "sandeep.com"
	fmt.Println(website)

	// no var style

	numberOfUser := 300000.0
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable LoginToken type: %T \n", LoginToken)

}
