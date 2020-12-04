package calculator_test

import (
	"calculator"
	"testing"
)
type addTest []struct {
	a float64 //input
	b float64 //input
	want float64 //output
	name string //label
}

type subtractTest []struct {
	a float64 // input
	b float64 //input
	want float64 //output
	name string //label
}

type multiplyTest []struct {
	a float64 //input
	b float64 //input
	want float64 //output
	name string //label
}

var addTests = addTest{
	{1, 1, 2, "Addition of two small integers"},
	{-10, -20, -30, "Addition of two negative integer"},
	{4.5, 4.5, 9, "Addition to two floating points"},
}

var subtractTests = subtractTest{
	{ 5, -5, 0, "Subtraction of positive and a negative integer"},
	{1, 1, 0, "Subtraction of two positive integers"},
	{10.5, 1.0, 9.5, "Subtraction of two floating points"},
}

var multiplyTests = multiplyTest{
	{4, 12, 48, "Multiplication of two positive integers"},
	{4, -12, -48, "Multiplication of a positive and negative integer"},
	{.5, 1, .5, "Multiplication of a floating point and an integer"},
}

func TestAdd(t *testing.T) {
	t.Parallel()
	for _, tt := range addTests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculator.Add(tt.a, tt.b)
		if tt.want != got {
			t.Errorf("%s: want %f, got %f",tt.name, tt.want, got)
		}
	})
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	for _, tt := range subtractTests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculator.Subtract(tt.a, tt.b)
			if tt.want != got {
				t.Errorf("%s: want %f, got %f", tt.name, tt.want, got)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	for _, tt := range multiplyTests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculator.Multiply(tt.a, tt.b)
			if tt.want != got {
				t.Errorf("%s: want %f, got %f", tt.name, tt.want, got)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	for _, tt := range multiplyTests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculator.Divide(tt.a, tt.b)
			if tt.want != got {
				t.Errorf("%s: want %f, got %f", tt.name, tt.want, got)
			}
		})
	}
}
