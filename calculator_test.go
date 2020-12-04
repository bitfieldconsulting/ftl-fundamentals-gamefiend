package calculator_test

import (
	"calculator"
	"testing"
)

func TestAddSubtractMultiply(t *testing.T) {
	t.Parallel()
	tcs := []struct {
		a    float64 // input
		b    float64 // input
		want float64 // output
		name string  // label
		f    func(float64, float64) float64
	}{
		{
			name: "Addition of two small integers",
			a:    1, b: 1, want: 2,
			f: calculator.Add,
		},
		{
			name: "Addition of two negative integer",
			a:    -10, b: -20, want: -30,
			f: calculator.Add,
		},
		{
			name: "Addition of two floating points",
			a:    4.5, b: 4.5, want: 9,
			f: calculator.Add,
		},
		{
			name: "Subtraction of positive and a negative integer",
			a:    5, b: -5, want: 10,
			f: calculator.Subtract,
		},
		{
			name: "Subtraction of two positive integers",
			a:    1, b: 1, want: 0,
			f: calculator.Subtract,
		},
		{
			name: "Subtraction of two floating points",
			a:    10.5, b: 1.0, want: 9.5,
			f: calculator.Subtract,
		},
		{
			name: "Multiplication of two positive integers",
			a:    4, b: 12, want: 48,
			f: calculator.Multiply,
		},
		{
			name: "Multiplication of a positive and negative integer",
			a:    0.5, b: 1.0, want: 0.5,
			f: calculator.Multiply,
		},
		{	name: "Multiplication of a floating point and an integer of two floating points",
			a:    10.5, b: 1.0, want: 10.5,
			f: calculator.Multiply,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.f(tc.a, tc.b)
			if tc.want != got {
				t.Errorf("want %f, got %f", tc.want, got)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	tcs := []struct {
		name        string
		a, b        float64
		want        float64
		errExpected bool
	}{
		{name: "Happy", a: 2, b: 1, want: 2},
		{name: "Zero divided by...", a: 0, b: 1, want: 0},
		{name: "Division by zero", a: 1, b: 0, want: 999, errExpected: true},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got, err := calculator.Divide(tc.a, tc.b)
			errReceived := err != nil
			if tc.errExpected != errReceived {
				t.Fatalf("unexpected error status %v", err)
			}
			if !tc.errExpected && tc.want != got {
				t.Errorf("%s: want %f, got %f", tc.name, tc.want, got)
			}
		})
	}
}
