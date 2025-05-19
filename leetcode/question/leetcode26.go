package question

func RemoveDuplicates(nums []int) int {
	fast, slow := 1, 0
	for fast < len(nums) && len(nums) >= 2 {
		if nums[slow] == nums[fast] {
			fast++
			continue
		}
		nums[slow+1] = nums[fast]
		slow++
		fast++
	}
	return slow + 1
}
