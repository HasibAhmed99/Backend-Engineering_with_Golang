package main

import (
	"fmt"
	"scope/furtherexplored"
)

// the imported package "fmt"
// is in the "file block" scope
// of this file

// x is in "package block" scope
	var x = 42
	func main () {
		//  x can access here
		fmt.Println(x)

		//  x can access here
		printMe()

		// y doesnt exist here yet
		// so this wont work
		// fmt.Println(y)

		// y is in "block scope"
		y:= 43 
		fmt.Println(y)

		p1 := person {
			"James",
		}
		p1.sayHello()

		// variable "shadowing"
		x := 32 
		fmt.Println(x)  

		// package block x still the same
		printMe() 

    furtherexplored.Fascinating()

		// also good to know
		if z := 82; z > 50 {
			fmt.Println(z)
		}
		// I cant access z here
		// fmt.Println(z)


		
	} 

	func printMe() {
		 	// x can be access here
			fmt.Println(x)
	}

	// type person is in "package block" scope
	type person struct {
		first string
	}

	// the method sayHello
	// which is attached to VALUES of TYPE person
	// is in "package block" scope
	func (p person) sayHello() {
			fmt.Println("Hi, my name is", p.first)
	}
	
	
               
