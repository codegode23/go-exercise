package main

import "fmt"

type Calculation func(int, int) int

func main() {
	var number_1 = 6
	var number_2 = 4

	var Addition Calculation = func(x, y int) int {
		return x + y
	}

	var Subtraction Calculation = func(x, y int) int {
		return x - y
	}

	fmt.Printf(
		"Addition of %v and %v: %v\n",
		number_1,
		number_2,
		Addition(number_1, number_2),
	)

	fmt.Printf(
		"Subtration of %v and %v: %v\n",
		number_1,
		number_2,
		Subtraction(number_1, number_2),
	)

}
