// Write a program which can compute the factorial of a given numbers.
//The results should be printed in a comma-separated sequence on a single line.

// Suppose the following input is supplied to the program: 8

// Then, the output should be: 40320

package main

import (
	"fmt"
)

func main() {

	var num, i, fact int64

	num = 8
	fact = 1

	for i = 1; i <= num; i++ {
		fact = fact * i
	}

	// Printing the factorial of the respective number
	fmt.Println("The factorial of", num, "is", fact)

	//Recursively
	// declaring the integer number using the var keyword
	// whose factorial we have to find
	var number int64

	// initializing the variable whose factorial we want to find
	number = 10

	// calling the factorial() function and printing the factorial
	fmt.Println("(Finding the factorial in a recursive manner.)")
	fmt.Println("The factorial of", number, "is", factorial(number))
}

//Recursively

func factorial(number int64) int64 {

	// if the number has reached 1 then we have to
	// return 1 as 1 is the minimum value we have to multiply with
	if number == 1 {
		return 1
	}

	// multiplying with the current number and calling the function
	// for 1 lesser number
	factorialOfNumber := number * factorial(number-1)

	// return the factorial of the current number
	return factorialOfNumber
}
