package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(10)
    y := rand.Intn(10)
	fmt.Printf(" x and y are %v and %v\t", x, y)

if x < 4 && y < 4 {
	fmt.Println("both are less than 4")
} else if x > 6 && y > 6 {
	fmt.Println("both are greater than 6")
} else if x >= 4 && x <= 6 {
	fmt.Println("x is from 4 to 6 inclusive of both numbers")
} else if y != 5 {
	fmt.Println("y is not five")
} else {
	fmt.Println("none of the previous met")
}
}