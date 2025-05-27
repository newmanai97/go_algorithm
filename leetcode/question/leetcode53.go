package question

import "math"

func maxSubArray(nums []int) int {
	result := math.MinInt
	count := 0

	for i := 0; i < len(nums); i++ {
		count = count + nums[i]
		if count > result {
			result = count
		}
		if count <= 0 {
			count = 0
		}
	}
	return result
}
