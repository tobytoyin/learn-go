package main

import (
	"fmt"
	"time"
	"strconv"
)

func pinger(s *string) {
	// infinite loop
	for i := 0; i < 5; i++ {
		*s = strconv.Itoa(i) // modify shared variable
		fmt.Printf("ping(%d)\n", i)
	}
}

func printer(s *string) {
	// retrieve values that are modified by `pinger`
	for {
		value := *s
		fmt.Printf("ref(%s)\n", value)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	// using Ptr instead of a channel
	// the value of ping-printer is not in sync 
	myValue := new(string)

	go pinger(myValue)
	go printer(myValue)

	// ref(0)
	// ping(0)
	// ping(1)
	// ping(2)
	// ping(3)
	// ping(4)
	// ref(4)
	// ref(4)
	// ref(4)	

	fmt.Scanln()
}
