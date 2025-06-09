package question

func fib1(n int) int {
	if n <= 1 {
		return n
	}
	return fib1(n-1) + fib1(n-2)
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	a, b, c := 0, 1, 0
	for i := 1; i < n; i++ {
		c = a + b
		a, b = b, c
	}
	return c
}
