package main

import "fmt"

func average(xs []float64) (avg float64) {
	total := 0.0

	for _, value := range xs {
		total += value
	}

	avg = total / float64(len(xs))
	return
}

func add(xs ...int) (sum int) {
	sum = 0

	for _, value := range xs {
		sum += value
	}
	return
}

func main() {
	slice1 := []float64{1, 2, 3}
	avg := average(slice1)

	fmt.Println(avg)

	fmt.Println(add(1, 2, 3))
}
