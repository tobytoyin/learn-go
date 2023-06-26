
package main

import "math"

// interface has an abstractmethod `area()`
type Shape interface {
	area() float64
}


// another function that sum all areas of any shapes
// by accessing the `Shape` interface
func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

// as long as other struct contains the same method
// this doesn't require any explicit extension
type Circle struct {
	radius float64
}

// Circle `area` method
func (circle *Circle) area() float64 {
	return math.Pi * math.Pow(2, circle.radius)
}

type Square struct {
	length float64
}

// Square `area` method
func (square *Square) area() float64 {
	return math.Pow(2, square.length)
}

func main() {
circle1 := &Circle{5.}
square1 := &Square{10.}

totalArea := totalArea(circle1, square1)

println("sum of area =", totalArea)


}