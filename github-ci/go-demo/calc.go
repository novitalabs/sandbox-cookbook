// Package main provides a tiny calculator demo used to exercise CI.
package main

import (
	"errors"
	"fmt"
)

// ErrDivideByZero is returned when dividing by zero.
var ErrDivideByZero = errors.New("divide by zero")

// Add returns the sum of a and b.
func Add(a, b int) int { return a + b }

// Sub returns a minus b.
func Sub(a, b int) int { return a - b }

// Mul returns the product of a and b.
func Mul(a, b int) int { return a * b }

// Div returns a divided by b, or ErrDivideByZero when b is zero.
func Div(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}

func main() {
	fmt.Println("2 + 3 =", Add(2, 3))
	fmt.Println("7 - 4 =", Sub(7, 4))
	fmt.Println("6 * 7 =", Mul(6, 7))
	if q, err := Div(10, 2); err == nil {
		fmt.Println("10 / 2 =", q)
	}
}
