package question

import (
	"math"
)

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if int(math.Abs(float64(target))) > sum {
		return 0
	}

	if (sum+target)%2 == 1 {
		return 0
	}

	bagweight := (sum + target) / 2
	dp := make([]int, bagweight+1)
	dp[0] = 1

	for i := 0; i < len(nums); i++ {
		for j := bagweight; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}
	return dp[bagweight]
}
