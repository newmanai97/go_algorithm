package question

func RemoveElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	i, j := 0, len(nums)-1

	for i <= j {
		if nums[i] != val {
			i++
			continue
		}
		if nums[j] == val {
			j--
			continue
		}
		if i == j {
			return i
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	return i
}

func RemoveElemen2(nums []int, val int) int {

	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		nums[slow] = nums[fast]
		if nums[slow] == val {
			continue
		}
		slow++
	}
	return slow
}
