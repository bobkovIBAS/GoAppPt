package server

type Calculator interface {
	FibonacciCalculation(n int) int
}

type FibonacciCalculatorData struct{}

func (f FibonacciCalculatorData) FibonacciCalculation(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}
