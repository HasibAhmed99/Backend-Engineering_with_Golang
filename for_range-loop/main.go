package main

import "fmt"

func main() {

	// for range loop
	xi := []int{41, 42, 43, 44, 45}
	for i, v := range xi {
		fmt.Println("ranging over a slice", i, v)
	}

	m := map[string]int{
		"James":      42,
		"Moneypenny": 43,
	}
	for k, v := range m {
		fmt.Println("ranging over a map", k, v)
	}
}
