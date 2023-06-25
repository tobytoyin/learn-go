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


// a function that interact with struct
// usuallly involves mutate some internal states of the struct 
// instead of return a new copy of it
// Thus, function taking struct as inputs often accepts struct's Pointer

func cirleArea(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

func main()  {
	// initialse struct
	c0 := new(Circle)  // using default values
	c1 := Circle{x:100, y:100, r:50}  // using input values


	fmt.Println(c0)  // printout the instance
	fmt.Println(c1.x, c1.y)  // printout the instance's fields

	fmt.Printf("area of c0 circle is %f \n", cirleArea(&c1))
}

