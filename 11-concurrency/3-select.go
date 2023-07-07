package main 

import (
	"time"
	"fmt"
)

type API struct {
	endpointA chan string
	endpointB chan string
}

func (s *API) Listener() {
	// `select` operates similar to `switch`
	// but it is switching base on the channel it receives
	// instead of the logical result of the variable

	// for example, 
	// like an API request, where one endpoint is using one channel, etc.
	// we can `select` different a handler to opreates on different channel for different endpoints
	for {
		select {
		case payloadA := <- s.endpointA:
			fmt.Println("callback HandlerA() on", payloadA)

		case payloadB := <- s.endpointB:
			fmt.Println("callback HandlerB() on", payloadB)
			
		case <- time.After(time.Second):
			fmt.Println("timeout")

		}
	}
} 


func main() {
	// we are creating two channels, i.e., variables
	// and store it inside a struct for better organisation
	app := new(API)
	app.endpointA = make(chan string)
	app.endpointB = make(chan string)

	go ModifyChannel1(app.endpointA)
	go ModifyChannel2(app.endpointB)
	
	// a background API listen to invoke handler
	go app.Listener()

	fmt.Scanln()

}

func ModifyChannel1(channel chan<- string) {
	for {
		channel <- "payload A"
		time.Sleep(time.Second * 2)
	}
}

func ModifyChannel2(channel chan<- string) {
	for {
		channel <- "payload B"
		time.Sleep(time.Second * 3)
	}
}