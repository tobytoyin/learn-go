package main

import (
	"fmt"
	"time"
	"strconv"
)

func pinger(channel chan string) {
	// infinite loop
	for i := 0; i < 5; i++ {
		channel <- strconv.Itoa(i)  // modify channel variable
		fmt.Printf("ping(%d)\n", i)		
	}
}

func printer(channel chan string) {
	for {
		value := <- channel
		fmt.Printf("ref(%s)\n", value)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	// channel variable works the same way as ordinary variable
	// but ensure the execution are sync, i.e., ping->printer; ping-> printer
	var channel chan string = make(chan string)

	go pinger(channel)
	go printer(channel)

	// ping(0)
	// ref(0)
	// ref(1)
	// ping(1)
	// ref(2)
	// ping(2)
	// ref(3)
	// ping(3)
	// ref(4)
	// ping(4)

	fmt.Scanln()
}
