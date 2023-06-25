package main

import (
	"fmt"
)

func main() {
	// closure means we can create a function within a function
	// in here: add func within main func

	// assign an lambda function and named it "add"
	add := func(x, y int) int {
		return x + y
	}

	fmt.Println(add(1, 2))

	// a inner function in Go cannot be declared using normal func syntax

	// func subtract(x, y int) int {
	// 	return x - y
	// }

}
