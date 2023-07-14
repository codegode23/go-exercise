package main

import "fmt"

type DivideError struct {
	Divident float64
}

func (de DivideError) Error() string {
	return fmt.Sprintf("can not divide %v by zero", de.Divident)
}

func divide(divident, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0.0, DivideError{Divident: divident}
	}

	return divident / divisor, nil
}

func main() {
	answer, err := divide(12, 3)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Answer:", answer)
	}

	answer, err = divide(24, 0)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Answer:", answer)
	}
}
