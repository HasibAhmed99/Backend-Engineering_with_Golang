package main

import (
	"fmt"

	"math/rand"
)

func main() {
	x := rand.Intn(400)
	fmt.Printf("The value of x is %v\n", x)

	if x <= 100 { 
	fmt.Println("between 1 to 100") 
} else if x >= 101 && x <= 200 {
	fmt.Println("between 101 to 200")
} else if x >= 201 && x <= 250 {
	fmt.Println("between 201 to 250")
} else {
	fmt.Println("over 250")
} 

fmt.Println(rand.Intn(3))
fmt.Println(rand.Intn(3))
fmt.Println(rand.Intn(3))
fmt.Println(rand.Intn(3))
fmt.Println(rand.Intn(3))
fmt.Println(rand.Intn(3))
fmt.Println(rand.Intn(3))
fmt.Println(rand.Intn(3))
fmt.Println(rand.Intn(3))
fmt.Println(rand.Intn(3))
fmt.Println(rand.Intn(3))
fmt.Println(rand.Intn(3))
fmt.Println(rand.Intn(3))
fmt.Println(rand.Intn(3))

}