// Write a program which accepts a sequence of comma-separated numbers
// from console and generate an slice out of them. Return the slice.

// Suppose the following input is supplied to the program: 34, 67, 55, 33, 12, 98.

// Then, the output should be: [34 67 55 33 12 98]

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	nums := "34, 67, 55, 33, 12, 98" //Declare the numbers

	numbers := strings.Split(nums, ",")

	length_of_numbers := len(numbers)
	var num = make([]int, length_of_numbers)

	for i, r := range numbers {
		str := strings.Trim(r, " ")
		num[i], _ = strconv.Atoi(str)
	}

	fmt.Println(num)
}
