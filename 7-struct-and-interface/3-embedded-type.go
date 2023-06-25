package main

import "fmt"

type Person struct {
	name string
}

func (p *Person) sayName() {
	println("My name is", p.name)
}

// embedded type shares the same concept as
// inheritance (subclass) in OOP
type Child struct {
	// a Child inherits as a Person
	Person
	school string
}

func main() {
	someFather := &Person{"timmy"}
	someFather.sayName()
	fmt.Printf("%T\n", someFather)

	someNamelessChild := new(Child)     // create struct with default value
	someNamelessChild.name = "john doe" // asign value later
	someNamelessChild.sayName()
	fmt.Printf("%T\n", someNamelessChild)

	someNamedChild := &Child{Person: Person{"timmy jr"}, school: "school"} // embedded another struct
	someNamedChild.sayName()
	fmt.Printf("%T\n", someNamedChild)
}
