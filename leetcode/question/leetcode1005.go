package question

import (
	"math"
	"sort"
)

func largestSumAfterKNegations(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return math.Abs(float64(nums[i])) > math.Abs(float64(nums[j]))
	})

	for i := 0; i < len(nums); i++ {
		if k > 0 && nums[i] < 0 {
			nums[i] = -nums[i]
			k--
		}
	}

	if k%2 == 1 {
		nums[len(nums)-1] = -nums[len(nums)-1]
	}
	sum := 0
	for j := 0; j < len(nums); j++ {
		sum = sum + nums[j]
	}
	return sum
}
