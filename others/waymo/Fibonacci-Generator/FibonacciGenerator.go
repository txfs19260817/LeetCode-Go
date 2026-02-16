package waymo

// Fib returns a stateful function that lazily produces Fibonacci numbers.
func Fib() func() int {
	a, b := 0, 1
	return func() int {
		v := a
		a, b = b, a+b
		return v
	}
}
