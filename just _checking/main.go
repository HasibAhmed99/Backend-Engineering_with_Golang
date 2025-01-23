package main

import "fmt"

var x = 42

func main() {
	/* 	switch {
	   	case x < 42:
	   		fmt.Println("x is LESS than 42")
	   	case x == 42:
	   		fmt.Println("x is EQUAL to 42")
	   	case x > 42:
	   		fmt.Println("x is GREATER than 42")
	   	default:
	   		fmt.Println("this is the default case for x")
	   	}*/
	/*   switch x {
	case 40:
		fmt.Println("x is 40")
	case 41:
		fmt.Println("x is 41")
	case 42:
		fmt.Println("x is 42")
	case 43:
		fmt.Println("x is 43")
	default:
		fmt.Println("this is the default case for x")
	}
}   */

// no default fallthrough
/* switch x {
case 42:
	fmt.Println("x is 42")
	fallthrough
case 43:
	fmt.Println("printed because of fallthrough: x is 43")
case 44:
	fmt.Println("printed because of fallthrough: x is 44")
case 45:
	fmt.Println("printed because of fallthrough: x is 45")
default:
	fmt.Println("printed because of fallthrough: this is the default case for x")
}
} */

// mo default fallthrough
switch x {

case 42:
	fmt.Println("x is 42")
	fallthrough
case 43:
	fmt.Println("printed because of fallthrough: x is 43")
	fallthrough
case 44:
	fmt.Println("printed because of fallthrough: x is 44")
	fallthrough
case 45:
    fmt.Println("printed because of fallthrough: x is 45")
	fallthrough
default:
	fmt.Println("printed because of fallthrough: this is the default case for x")
}
}
