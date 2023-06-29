package main

import (
	"errors"
	"fmt"
)

func OnlyAllowPositiveNumbers(x float32) {
	if x < 0.0 {
		err := errors.New("value is not positive")  // creating an error message
		panic(err)  // throwing error message
	}

	fmt.Println(x)
}

func main() {
	OnlyAllowPositiveNumbers(5.0)  // this would be acceptable
	OnlyAllowPositiveNumbers(-5.0)  // this would not throwing erro
}

