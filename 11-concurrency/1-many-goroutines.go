package main

import (
	"fmt"
	"time"
)

func PrintNumbers(id int) {
	for i := 0; i < 3; i++ {
		fmt.Printf("PrintNumbers(%d) - %d / 5\n", id, i)
	}
}

func BlockFunction() {
	fmt.Println("do something else.")
	time.Sleep(time.Second * 5)
	fmt.Println("end something else.")
}

func main() {
	// this are execute and NOT wait for the return result
	for id := 0; id < 5; id ++ {
		go PrintNumbers(id)
	}
	
	// add a blockingg function to wait for concurrency to complete
	BlockFunction()
}
