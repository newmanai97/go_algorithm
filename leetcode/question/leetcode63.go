package question

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	if obstacleGrid[0][0] == 1 || obstacleGrid[m-1][n-1] == 1 {
		return 0
	}

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		if obstacleGrid[i][0] == 1 {
			dp[i][0] = 0
			break
		} else {
			dp[i][0] = 1
		}
	}

	for j := 0; j < n; j++ {
		if obstacleGrid[0][j] == 1 {
			dp[0][j] = 0
			break
		} else {
			dp[0][j] = 1
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				continue
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}
