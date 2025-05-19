package question

import "sort"

var (
	path5 []int
	res5  [][]int
	used  []bool
)

func combinationSum2(candidates []int, target int) [][]int {
	path5, res5, used = make([]int, 0, len(candidates)), make([][]int, 0), make([]bool, len(candidates))
	sort.Ints(candidates)
	dfs5(candidates, 0, target, 0)
	return res5
}

func dfs5(candidates []int, start int, target int, sum int) {
	if sum == target {
		tmp := make([]int, len(path5))
		copy(tmp, path5)
		res5 = append(res5, tmp)
		return
	}

	for i := start; i <= len(candidates)-1; i++ {
		if sum > target {
			break
		}

		if i > 0 && candidates[i] == candidates[i-1] && used[i-1] == false {
			continue
		}
		path5 = append(path5, candidates[i])
		used[i] = true
		dfs5(candidates, i+1, target, sum+candidates[i])
		path5 = path5[:len(path5)-1]
		used[i] = false
	}
}
