package question

func climbStairs1(n int) int {
	if n <= 2 {
		return n
	}
	return climbStairs(n-1) + climbStairs(n-2)
}

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	a, b, c := 1, 2, 0
	for i := 2; i < n; i++ {
		c = a + b
		a, b = b, c
	}
	return c
}
