package main

import (
	"fmt"

	"math/rand"
)

func init() {
	fmt.Println("This from the init function")
}

func main() {
	x := rand.Intn(400)
	fmt.Printf("the value of x is: %v\n", x)

	switch {
	case x <= 100:
		fmt.Println("between 0 to 100")
	case x >= 101 && x <= 200:
		fmt.Println("between 101 to 200")
	case x >= 201 && x <= 250:
		fmt.Println("between 201 to 250")
	default:
		fmt.Println("over 250")	
	}
}