package question

func jump(nums []int) int {
	step := 0
	curcover := 0
	nextcover := 0
	if len(nums) == 1 {
		return step
	}
	for i := 0; i <= len(nums)-1; i++ {
		if i == curcover+1 {
			step++
			curcover = nextcover
		}
		if i+nums[i] > nextcover {
			nextcover = i + nums[i]
		}
	}
	return step
}
