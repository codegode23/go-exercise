package main

import (
	"fmt"
	"sort"
)

func main() {
	// Sorting Ints
	numbers := []int{42, 12, 67, 3, 99, 1}
	sort.Ints(numbers)
	fmt.Println("Sorted Ints:", numbers)

	// Sorting Strings
	names := []string{"Alice", "Eve", "Bob", "Charlie"}
	sort.Strings(names)
	fmt.Println("Sorted Strings:", names)

	// Sorting Float64s
	floats := []float64{3.14, 1.618, 2.718, 0.577}
	sort.Float64s(floats)
	fmt.Println("Sorted Float64s:", floats)
}
