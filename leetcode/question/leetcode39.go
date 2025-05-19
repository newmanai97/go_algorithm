package question

import "sort"

var (
	path4 []int
	res4  [][]int
)

func combinationSum(candidates []int, target int) [][]int {
	path4, res4 = make([]int, 0, len(candidates)), make([][]int, 0)
	sort.Ints(candidates)
	dfs4(candidates, 0, target, 0)
	return res4
}

func dfs4(candidates []int, start int, target int, sum int) {
	if sum == target {
		tmp := make([]int, len(path4))
		copy(tmp, path4)
		res4 = append(res4, tmp)
		return
	}
	for i := start; i <= len(candidates)-1; i++ {
		if sum+candidates[i] > target {
			break
		}
		path4 = append(path4, candidates[i])
		dfs4(candidates, i, target, sum+candidates[i])
		path4 = path4[:len(path4)-1]
	}
}
