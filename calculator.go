// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"errors"
	"fmt"
	"math"
)

// Add takes two numbers and returns the result of adding them together.
func Add(a, b float64, nums ...float64) float64 {
	total := a + b
	if len(nums) > 0 {
		for _, d := range nums {
			total += d
		}
	}
	return total
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(a, b float64, nums ...float64) float64 {
	total := a - b
	if len(nums) > 0 {
		for _, d := range nums {
			total -= d
		}
	}
	return total
}

// Multiply takes two numbers and returns the result of multiplying the first by
// the second.
func Multiply(a, b float64, nums ...float64) float64 {
	total := a * b
	if len(nums) > 0 {
		for _, d := range nums {
			total *= d
		}
	}
	return total

}

// Divide takes two numbers and returns the result of dividing the first by the
// second, or an error if the second value is zero.
func Divide(nums ...float64) (float64, error) {
	total := nums[0]
	for i, d := range nums {
		// skip the first element, which is the total
		if i == 0 {
			continue
		}
		// return error if divisor is 0
		if d == 0 {
			return 0, errors.New("division by zero is undefined")
		}
		total /= d
	}
	return total, nil
}

func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, fmt.Errorf("bad input %f: square root of negative number is undefined", a)
	}
	return math.Sqrt(a), nil
}
