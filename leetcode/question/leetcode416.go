package question

func canPartition(nums []int) bool {
	sum := 0

	for _, v := range nums {
		sum += v
	}

	if sum%2 == 1 {
		return false
	}

	target := sum / 2
	dp := make([]int, target+1)

	for _, v := range nums {
		for j := target; j >= v; j-- {
			if dp[j] < dp[j-v]+v {
				dp[j] = dp[j-v] + v
			}
		}
	}
	return dp[target] == target
}
