package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// selecting a channel
	ch1 := make(chan int)
	ch2 := make(chan int)

	d1 := time.Duration(rand.Int63n(250))
	d2 := time.Duration(rand.Int63n(250))
	fmt.Printf("%v \t %T\n", d1, d1)
	time.Sleep(d1 * time.Millisecond)
	fmt.Println("Done!")


go func() {
	time.Sleep(d1 * time.Millisecond)
	ch1 <- 41
}()

go func() {
	time.Sleep(d2 * time.Millisecond)
	ch2 <- 42
}()

   // A "select" statement chooses which of a set of possible send or receive operations will proceed.
   select {
   case v1 := <-ch1:
	fmt.Println("value from channel 1", v1)
   case v2 := <-ch2:
	fmt.Println("value from channel 2", v2)
   }
   }
   
