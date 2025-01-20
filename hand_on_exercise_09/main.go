package main

import "fmt"

var a = 40

const b = 41

func main() {
	c := 42
	fmt.Printf("the value of a is %v and the type of a is %T\n", a, a)
	fmt.Printf("the value of b is %v and the type of b is %T\n", b, b)
	fmt.Printf("the value of c is %v and the type of c is %T\n", c, c)
}