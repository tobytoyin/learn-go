package main

import "fmt"

// a closure function that returns:
// another function; that returns an int
func accumulator() func() uint {
	i := uint(0)

	_accumulator := func() (ret uint) {
		ret = i
		i++
		return ret
	}
	return _accumulator
}

func main() {
	accum := accumulator()

	fmt.Println(accum())
	fmt.Println(accum())
	fmt.Println(accum())
	fmt.Println(accum())

}
