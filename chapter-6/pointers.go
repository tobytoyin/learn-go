package main

import "fmt"

func setZero(xPtr *int) {
	// this funtion set a variable's reference to be 0
	*xPtr = 0
}

func main() {
	x := 5
	xPtr := &x // retrieve pointer from variables

	fmt.Printf("x = %d \n", x)
	setZero(xPtr)
	fmt.Printf("x = %d \n", x)

	yPtr := new(int) // create a new pointer immediately
	setZero(yPtr)
	y := *yPtr // retrieve the pointer's variable
	fmt.Printf("y = %d \n", y)

}
