package mymath

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	sum := Add(1.0, 2.0)
	if sum != 3.0 {
		t.Fatalf("Error in calculation")
	}
}

type TestCase struct {
	value1   float64
	value2   float64
	expected float64
}

//To run multiple cases
func TestMany(t *testing.T) {
	testCases := []TestCase{
		{10.0, 10.0, 20.0},
		{0.0, 0.0, 0.0},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%f + %f ", tc.value1, tc.value2), func(t *testing.T) {
			sum := Add(tc.value1, tc.value2)
			if sum != tc.expected {
				t.Fatalf("Error in calculation")
			}
		})
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		sum := Add(float64(i), float64(i+1))
		if sum == 0 {
			b.Fatal("Benchmarking error.")
		}
	}
}
