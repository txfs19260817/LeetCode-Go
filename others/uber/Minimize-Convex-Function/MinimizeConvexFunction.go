package uber

// ConvexFunction provides access to a black-box convex function.
type ConvexFunction interface {
	Evaluate(x float64) float64
}

// Solution exposes the minimization API over a convex function.
type Solution struct {
	fn ConvexFunction
}

// NewSolution constructs a Solution with the provided convex function.
func NewSolution(fn ConvexFunction) *Solution {
	return &Solution{fn: fn}
}

// Minimize finds a point x* in [a, b] such that |x* - x_min| <= eps.
func (s *Solution) Minimize(a float64, b float64, eps float64) float64 {
	for b-a > eps {
		xa, xb := a+(b-a)/3, b-(b-a)/3
		fa, fb := s.fn.Evaluate(xa), s.fn.Evaluate(xb)
		if fa < fb {
			b = xb
		} else {
			a = xa
		}
	}
	return (a + b) / 2
}
