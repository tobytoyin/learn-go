package main

import "fmt"

type Person struct {
	name string
	id   int
}

func main() {
	// initialising a struct
	var person1 Person
	person := new(Person)
	person2 := &Person{"Toby", 123}

	fmt.Println(person1)
	fmt.Println(person)
	fmt.Println(person2)

}
