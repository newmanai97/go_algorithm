package question

func isPerfectSquare(num int) bool {
	left, right := 0, num
	ans := false

	for left <= right {
		mid := left + (right-left)>>1
		if mid*mid == num {
			ans = true
			break
		} else if mid*mid > num {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}
