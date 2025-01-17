// package main

// import "fmt"

// func main() {
// 	fmt.Println("Hello World")
// }

/*  package main

import "fmt"
func main ()  {
	a := 20
	a =  30
	fmt.Println(a)
	a = 40
	var a const = 50
	a = 60 // it cant be change like var because its a constant which can not be change

} */

// [Golang] - 009 => if... else and switch
/*  package main

import "fmt"

func main () {
// 	age := 20
// 	sex := "female"

// 	if age == 20 || sex == "male" {
// 	fmt.Println("You are ready to be married!")
// } else {
// fmt.Println("You are not ready to be married!")
// 	}
isPretty := false

if !isPretty {
	fmt.Println("Print something.....")
}
// >, >=, <, =>, ==
// and => &&
// or => ||
// not => !
} */

/* package main

import "fmt"
func main ()  {
	a := 100

	switch a {
	case 1:
		fmt.Println("a is 1")
    case 2, 3:
		fmt.Println("a is either 2 or 3")
	default:
		fmt.Println("a is neither 1, 2 or 3")
		}
}	*/

//[Golang]- 010

/*  package main

import "fmt"

func add(num1 int, num2 int) {
sum := num1 + num2

fmt.Println(sum) }
func main ()  {
	a := 10
	b := 20
add(a, b)

add(5, 7)

 } 	*/

// [Golang] -> 011 : function with return value and types
/*	package main

import "fmt"

// func add(num1 int, num2 int) int {
// 	sum := num1 + num2

// 	return sum
// }


func getNumbers(num1 int, num2 int) (int, int) {
	sum := num1 + num2
	mul := num1 * num2

	return sum, mul
}

func main () {
	a := 10
	b := 20

	p, q := getNumbers(a, b)
	fmt.Println(p)
	fmt.Println(q)
}   */

// [Golang] -> 012: more function examples
/* package main

import "fmt"

func add(num1 int, num2 int) int {
	sum := num1 + num2

	return sum
}

func getNumbers(num1 int, num2 int) (int, int) {
	sum := num1 + num2
	mul := num1 * num2

	return sum, mul
}

func printSomething() {
	fmt.Println("Education must be free!")
}

func sayHello(name string) {
	fmt.Println("Welcome to the golang course,", name)
}

func main() {
	// a := "Habib"
	printSomething()
	sayHello("Habib")

} */

// [Golang]-013 => Why function is needed
/* 	package main

import "fmt"

func main() {
	//Print welcome message
	fmt.Println("Welcome to the application")

	// Get user as input
	var name string
	fmt.Println("Enter your name - ")
	fmt.Scanln(&name)

	// get numbers
	var num1 int
	var num2 int
	fmt.Println("Enter your first number - ")
	fmt.Scanln(&num1) // & => ampersand
	fmt.Println("Enter your second number - ")
	fmt.Scanln(&num2)
	sum := num1 + num2

	// Display results
	fmt.Println("Hello, ", name)
	fmt.Println("Summation = ", sum)

	// Print a goodbye message
	fmt.Println("Thank you for using the application")
	fmt.Println("Goodbye!")
}	*/

// But we can code those beatiful way which will easy to read

package main

import "fmt"

func printWelcomeMessge() {
	fmt.Println("Welcome to the application")
}

func getUserName() string {
	var name string
	fmt.Println("Enter your name - ")
	fmt.Scanln(&name)
	return name
}

func getNumbers() {
	var num1 int
	var num2 int
	fmt.Println("Enter your first number - ")
	fmt.Scanln(&num1) // & => ampersand
	fmt.Println("Enter your second number - ")
	fmt.Scanln(&num2)
	sum := num1 + num2 
	fmt.Println(sum)
}

func main() {
	//Print welcome message
	printWelcomeMessge()

	// Get user as input
getUserName()

	// get numbers 
	getNumbers()

	// Display results
	fmt.Println("Hello, ", name)
	fmt.Println("Summation = ", sum)

	// Print a goodbye message
	fmt.Println("Thank you for using the application")
	fmt.Println("Goodbye!")
}