package question

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i, _ := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < len(strs); i++ {
		zeronum, onenum := 0, 0
		for _, c := range strs[i] {
			if c == '0' {
				zeronum++
			} else {
				onenum++
			}
		}

		for j := m; j >= zeronum; j-- {
			for k := n; k >= onenum; k-- {
				if dp[j][k] < dp[j-zeronum][k-onenum]+1 {
					dp[j][k] = dp[j-zeronum][k-onenum] + 1
				}
			}
		}
	}
	return dp[m][n]
}
