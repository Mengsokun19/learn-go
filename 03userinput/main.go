package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "this is the simple for user input: "
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza: ")

	// comma ok syntax or erorr okay syntax (try catch in other languages)
	input, _ := reader.ReadString('\n') // read until the new line character or user click enter
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of this rating is %T ", input)
}
