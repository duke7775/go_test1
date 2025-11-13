package main

import (
	"fmt"
)

type DivideError struct {
	Dividend int
	Divisor  int
}

func (D *DivideError) Error() string {
	return fmt.Sprintf("cannot divide %d by %d", D.Dividend, D.Divisor)
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, &DivideError{Dividend: a, Divisor: b}
	}
	return a / b, nil
}

func main() {
	_, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
