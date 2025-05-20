package question

var (
	path11 []int
	res11  [][]int
	used11 []bool
)

func permute(nums []int) [][]int {
	path11 = make([]int, 0, len(nums))
	res11 = make([][]int, 0)
	used11 = make([]bool, len(nums))
	dfs11(nums, 0)
	return res11
}

func dfs11(nums []int, start int) {
	if start >= len(nums) {
		tmp := make([]int, len(path11))
		copy(tmp, path11)
		res11 = append(res11, tmp)
	}

	for i := 0; i < len(nums); i++ {
		if !used11[i] {
			path11 = append(path11, nums[i])
			used11[i] = true
			dfs11(nums, start+1)
			used11[i] = false
			path11 = path11[:len(path11)-1]
		}
	}

}
