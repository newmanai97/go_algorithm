package question

import (
	"reflect"
	"sort"
)

var (
	path9 []int
	res9  [][]int
)

func subsetsWithDup(nums []int) [][]int {
	path9, res9 = make([]int, 0), make([][]int, 0)
	sort.Ints(nums)
	dfs9(nums, 0)
	return res9
}

func dfs9(nums []int, start int) {
	sort.Ints(path9)
	for _, y := range res9 {
		if reflect.DeepEqual(y, path9) {
			return
		}
	}
	tmp := make([]int, len(path9))
	copy(tmp, path9)
	res9 = append(res9, tmp)
	if start >= len(nums) {
		return
	}

	for i := start; i < len(nums); i++ {
		path9 = append(path9, nums[i])
		dfs9(nums, i+1)
		path9 = path9[:len(path9)-1]
	}
}
