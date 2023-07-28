package main

import (
	"fmt"
	"math/big"

	// "math/rand"
	"crypto/rand"
)

func main() {
	fmt.Println("Welcom to math in Go lang")

	var myNumberOne int = 2
	var myNumberTwo float64 = 5

	fmt.Println("The sum of the 2 numbers is: ", myNumberOne+int(myNumberTwo))

	// random number using math
	// fmt.Println("This is the random number:", rand.Intn(10))

	// random number using crypto
	cryptoRandomNumber, _ := rand.Int(rand.Reader, big.NewInt(10))
	fmt.Println("This is the random number from crypto package:", cryptoRandomNumber)
}
