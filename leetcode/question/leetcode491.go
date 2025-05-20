package question

var (
	path10 []int
	res10  [][]int
)

func findSubsequences(nums []int) [][]int {
	path10, res10 = make([]int, 0), make([][]int, 0)
	dfs10(nums, 0)
	return res10
}

func dfs10(nums []int, start int) {
	if len(path10) > 1 {
		tmp := make([]int, len(path10))
		copy(tmp, path10)
		res10 = append(res10, tmp)
	}

	set := make(map[int]interface{})
	for i := start; i < len(nums); i++ {
		if _, exist := set[nums[i]]; exist {
			continue
		}
		if len(path10) >= 1 && nums[i] < path10[len(path10)-1] {
			continue
		}

		path10 = append(path10, nums[i])
		set[nums[i]] = nil
		dfs10(nums, i+1)
		path10 = path10[:len(path10)-1]
	}
}
