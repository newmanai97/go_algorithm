package question

import "sort"

var (
	path12 []int
	res12  [][]int
	used12 []bool
)

func permuteUnique(nums []int) [][]int {
	path12 = make([]int, 0, len(nums))
	res12 = make([][]int, 0)
	used12 = make([]bool, len(nums))
	sort.Ints(nums)
	dfs12(nums, 0)
	return res12

}

func dfs12(nums []int, cur int) {
	if len(path12) == len(nums) {
		tmp := make([]int, len(path12))
		copy(tmp, path12)
		res12 = append(res12, tmp)
	}

	for i := 0; i < len(nums); i++ {
		if !used12[i] {
			if i >= 1 && nums[i] == nums[i-1] && used12[i-1] == false {
				continue
			}
			path12 = append(path12, nums[i])
			used12[i] = true
			dfs12(nums, cur+1)
			used12[i] = false
			path12 = path12[:len(path12)-1]
		}
	}
}
