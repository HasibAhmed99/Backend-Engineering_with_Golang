package main

import (
	"fmt"

	"github.com/HasibAhmed99/puppy"
)

func main() {
	s1 := puppy.Bark()
	s2 := puppy.Barks()
	fmt.Println(s1)
	fmt.Println(s2)

	// i can also do like this
	fmt.Println(puppy.Bark())
	fmt.Println(puppy.Barks())
}
