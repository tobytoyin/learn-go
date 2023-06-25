package main

import "fmt"

// defining a struct (eqv class) this way
type Person struct {
	name  string
	age   int
	score float32
}

func (person *Person) printName() {
	fmt.Printf("name is %s\n", person.name)
}

func main() {
	myPerson := Person{"timmy", 30, 80.0} // initialise with input args
	myPerson.printName()
}
