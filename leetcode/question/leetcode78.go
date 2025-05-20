package question

var (
	path8 []int
	res8  [][]int
)

func subsets(nums []int) [][]int {
	path8, res8 = make([]int, 0), make([][]int, 0)
	dfs8(nums, 0)
	return res8
}

func dfs8(nums []int, start int) {
	tmp := make([]int, len(path8))
	copy(tmp, path8)
	res8 = append(res8, tmp)
	if start >= len(nums) {
		return
	}

	for i := start; i < len(nums); i++ {
		path8 = append(path8, nums[i])
		dfs8(nums, i+1)
		path8 = path8[:len(path8)-1]
	}
}
