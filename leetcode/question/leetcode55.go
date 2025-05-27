package question

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	cover := 0
	for i := 0; i <= cover; i++ {
		if i+nums[i] > cover {
			cover = i + nums[i]
		}
		if cover >= len(nums)-1 {
			return true
		}
	}
	return false
}
