package uber

import (
	"math"
	"testing"
)

type testConvexFunction struct {
	f func(float64) float64
}

func (t testConvexFunction) Evaluate(x float64) float64 {
	return t.f(x)
}

func TestMinimizeConvexFunction(t *testing.T) {
	tests := []struct {
		name     string
		f        func(float64) float64
		a        float64
		b        float64
		eps      float64
		expected float64
	}{
		{
			name:     "Quadratic centered at 3",
			f:        func(x float64) float64 { return (x-3)*(x-3) + 5 },
			a:        -10,
			b:        10,
			eps:      0.01,
			expected: 3.0,
		},
		{
			name:     "Quadratic centered at 0",
			f:        func(x float64) float64 { return x * x },
			a:        -100,
			b:        50,
			eps:      0.001,
			expected: 0.0,
		},
		{
			name:     "Absolute value",
			f:        func(x float64) float64 { return math.Abs(x - 123.456) },
			a:        0,
			b:        1000,
			eps:      0.0001,
			expected: 123.456,
		},
		{
			name:     "Large interval",
			f:        func(x float64) float64 { return (x - 987.654) * (x - 987.654) },
			a:        -1e9,
			b:        1e9,
			eps:      1e-4,
			expected: 987.654,
		},
		{
			name:     "Exponential minus linear",
			f:        func(x float64) float64 { return math.Exp(x) - 2*x },
			a:        0,
			b:        2,
			eps:      1e-4,
			expected: 0.693147,
		},
		{
			name:     "Linear function minimum at boundary",
			f:        func(x float64) float64 { return x },
			a:        0,
			b:        100,
			eps:      1e-4,
			expected: 0.0,
		},
		{
			name:     "Tight interval around minimum",
			f:        func(x float64) float64 { return (x - 1) * (x - 1) },
			a:        0.9999,
			b:        1.0001,
			eps:      1e-4,
			expected: 1.0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := &Solution{fn: testConvexFunction{f: tc.f}}
			result := s.Minimize(tc.a, tc.b, tc.eps)

			if result < tc.a || result > tc.b {
				t.Fatalf("Minimize returned %v outside [%v, %v]", result, tc.a, tc.b)
			}

			if math.Abs(result-tc.expected) > tc.eps {
				t.Fatalf("Minimize(%v, %v, %v) = %v; want within %v of %v", tc.a, tc.b, tc.eps, result, tc.eps, tc.expected)
			}
		})
	}
}
