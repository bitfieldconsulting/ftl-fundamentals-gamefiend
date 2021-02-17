package calculator_test

import (
	"calculator"
	"testing"
)

func TestAddSubtractMultiply(t *testing.T) {
	t.Parallel()
	tcs := []struct {
		nums []float64
		want float64 // output
		name string  // label
		f    func(a, b float64, nums ...float64) float64
	}{
		{
			name: "Addition of two small integers",
			nums: []float64{1, 1},
			want: 2,
			f:    calculator.Add,
		},
		{
			name: "Addition of two negative integer",
			nums: []float64{-10, -20},
			want: -30,
			f:    calculator.Add,
		},
		{
			name: "Addition of two floating points",
			nums: []float64{4.5, 4.5},
			want: 9,
			f:    calculator.Add,
		},
		{
			name: "Addition of four floating points",
			nums: []float64{4.5, 4.5, 4.5, 4.5},
			want: 18,
			f:    calculator.Add,
		},
		{
			name: "Subtraction of positive and a negative integer",
			nums: []float64{5, -5},
			want: 10,
			f:    calculator.Subtract,
		},
		{
			name: "Subtraction of two positive integers",
			nums: []float64{1, 1},
			want: 0,
			f:    calculator.Subtract,
		},
		{
			name: "Subtraction of two floating points",
			nums: []float64{10.5, 1.0},
			want: 9.5,
			f:    calculator.Subtract,
		},
		{
			name: "Multiplication of two positive integers",
			nums: []float64{4, 12},
			want: 48,
			f:    calculator.Multiply,
		},
		{
			name: "Multiplication of a positive and negative integer",
			nums: []float64{0.5, 1.0},
			want: 0.5,
			f:    calculator.Multiply,
		},
		{
			name: "Multiplication of two floating points",
			nums: []float64{10.5, 1.0},
			want: 10.5,
			f:    calculator.Multiply,
		},
		{
			name: "Multiplication of three floating points",
			nums: []float64{10.5, 1.0, 3.0},
			want: 31.5,
			f:    calculator.Multiply,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.f(tc.nums[0], tc.nums[1], tc.nums[2:]...)
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
		nums        []float64
		want        float64
		errExpected bool
	}{
		{name: "Happy", nums: []float64{2, 1}, want: 2},
		{name: "Happy with three divisors", nums: []float64{20, 5, 2}, want: 2},
		{name: "Zero divided by...", nums: []float64{0, 20}, want: 0},
		{name: "Zero divided by...with three divisors", nums: []float64{0, 10, 20}, want: 0},
		{name: "Division by zero", nums: []float64{1, 0}, want: 999, errExpected: true},
		{name: "Division by zero with three divisors", nums: []float64{10, 5, 0}, want: 999, errExpected: true},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got, err := calculator.Divide(tc.nums...)
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

func TestSqrt(t *testing.T) {
	t.Parallel()
	tcs := []struct {
		name        string
		input       float64
		want        float64
		errExpected bool
	}{
		{name: "Happy", input: 4, want: 2},
		{name: "Zero", input: 0, want: 0},
		{name: "Negative", input: -1, want: 999, errExpected: false},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got, err := calculator.Sqrt(tc.input)
			errReceived := err != nil
			if tc.errExpected != errReceived {
				t.Fatalf("unexpected error status: %v", err)
			}
			if !tc.errExpected && tc.want != got {
				t.Errorf("%s: want %f, got %f", tc.name, tc.want, got)
			}
		})
	}
}
