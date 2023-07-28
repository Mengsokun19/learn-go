package main

import "fmt"

// jwtToken := 100.123 // error
var jwtToken = 123.123 // global variable

// Constant
const LoginToken string = "asdasdasdasdasdasdasdasd"

func main() {
	fmt.Println("public variable: ", LoginToken)

	// * String
	var user string
	user = "John Doe"
	fmt.Println("User name is: ", user)
	fmt.Printf("Variable is of type: %T \n", user)

	// * Boolean
	var isLoggedIn bool
	isLoggedIn = true
	fmt.Println("Is user logged in? ", isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	// * Integer
	var smallVal uint16 = 256
	fmt.Println("Small value: ", smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	// * Float
	var smallFloat float32 = 255.5
	fmt.Println("Small float value: ", smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// ? Default values and some aliases
	var anotherVariable int
	fmt.Println("Default value of int: ", anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// Implicit type
	var wesbite = "youtube.com"
	fmt.Println("Website: ", wesbite)

	// No var style declaration
	numberOfPeopel := 100.123
	fmt.Println("Number of people: ", numberOfPeopel)
}
