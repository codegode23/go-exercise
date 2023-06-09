// Write a program which accepts a sequence of comma-separated numbers
// from console and generate an slice out of them. Return the slice.

// Suppose the following input is supplied to the program: 34, 67, 55, 33, 12, 98.

// Then, the output should be: [34 67 55 33 12 98]

package main

import "fmt"

func main() {
	// String as a range in the for loop
	for index, s := range "GeeksForGeeKs" {

		fmt.Printf("The index number of %c is %d\n", s, index)
	}
}
