package main

import ("fmt"; "math")

// struct is similar to class that stores different related variables
// into a single data structure

type Circle struct {
	//	fields of a struct, i.e., attributes
	x float64
	y float64
	r float64
}


// method for struct
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}


func main()  {
	// initialse struct
	c0 := Circle{x:100, y:100, r:50} 


	fmt.Println(c0)  // printout the instance
	fmt.Printf("area of c0 circle is %f \n", c0.area())
}

