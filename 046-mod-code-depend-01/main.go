package main

import (
	"fmt"

	"github.com/HasibAhmed99/tommy"
)

func main() {
   s1:= tommy.Bark()
   s2:= tommy.Barks()
   fmt.Println(s1)
   fmt.Println(s2)

   // also like this
   fmt.Println(tommy.Bark())
   fmt.Println(tommy.Barks()) 
    
}