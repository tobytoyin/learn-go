package main

import (
	"fmt"
	"time"
)

func PrintNumbers(n int) {
	for i := 0; i < n+1; i++ {
		fmt.Println(i, "/", n)
	}
}

func BlockFunction() {
	fmt.Println("do something else.")
	time.Sleep(time.Second * 5)
	fmt.Println("end something else.")
}

func main() {
	// this are execute and NOT wait for the return result
	go PrintNumbers(5)

	// add a blockingg function to wait for concurrency to complete
	BlockFunction()
}
