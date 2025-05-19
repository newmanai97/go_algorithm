package question

import "math"

func minSubArrayLen(target int, nums []int) int {
	result := math.MaxInt
	sum, i, j, sublength := 0, 0, 0, 0
	for j < len(nums) {
		sum += nums[j]
		for sum >= target {
			sublength = j - i + 1
			if sublength < result {
				result = sublength
			}
			sum -= nums[i]
			i++
		}
		j++
	}
	if result == math.MaxInt {
		return 0
	}
	return result
}
