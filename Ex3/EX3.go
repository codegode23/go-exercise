// With a given integral number n, write a program to generate a
// map that contains (i, i*i) such that is an integral number
// between 1 and n (both included), and then the program should
// print the map with representation of the value

// Suppose the following input is supplied to the program: 8

// Then, the output should be: map[1:1 2:4 3:9 4:16 5:25 6:36 7:49 8:64]
// Go program to illustrate how to
// create and initialize maps
package main

import "fmt"

func main() {

	var n int //here n is the size of the created map

	var numbers = make(map[int]int, n)

	n = 8

	for i := 1; i <= n; i++ {

		numbers[i] = i * i
	}

	fmt.Println(numbers)

}
