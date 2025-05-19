package question

var (
	path2 []int
	res   [][]int
)

func combinationSum3(k int, n int) [][]int {
	path2, res = make([]int, k), make([][]int, 0)
	dfs2(k, n, 1, 0)
	return res
}

func dfs2(k, n, start, sum int) {
	if len(path2) == k {
		if sum == n {
			tmp := make([]int, k)
			copy(tmp, path2)
			res = append(res, path2)
		}
	}

	for i := start; i <= 9; i++ {
		if sum+i > n || 9-i+1 < k-len(path2) {
			break
		}
		path2 = append(path2, i)
		dfs2(k, n, i+1, sum+i)
		path2 = path2[:len(path2)-1]
	}
}
