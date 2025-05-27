package question

func wiggleMaxLength(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}
	ans := 1
	prediff := nums[1] - nums[0]
	if prediff != 0 {
		ans = 2
	}
	for i := 2; i < n; i++ {
		curdiff := nums[i] - nums[i-1]
		if (curdiff > 0 && prediff <= 0) || (curdiff < 0 && prediff >= 0) {
			ans++
			prediff = curdiff
		}

	}
	return ans
}
