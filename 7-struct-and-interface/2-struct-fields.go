package main

import "fmt"

// defining a struct (eqv class) this way
type Person struct {
	name string
	age int
	score float32
}

// this creates a `Person` type struct-object
// with 3 fields `name`, `age`, `score`

func main() { 
	// initialisation involving Copy of Object
	var defaultPerson Person  // initialise with default args
	fmt.Println(defaultPerson) 

	myPerson := Person{"timmy", 30, 80.0}  // initialise with input args
	fmt.Println(myPerson)

	// initialisation invovling Pointers
	defaultPersonPtr := new(Person)  // initialise with default args
	fmt.Println(defaultPersonPtr)

	myPersonPtr := &Person{"tommy", 40, 90.0}  // initialise with input args
	fmt.Println(myPersonPtr)
}