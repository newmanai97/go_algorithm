package question

func SearchRange(nums []int, target int) []int {
	i, j := 0, len(nums)-1
	k := -1
	for i <= j {
		mid := i + (j-i)/2
		if nums[mid] == target {
			k = mid
		} else if nums[mid] > target {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}
	if k == -1 {
		return []int{-1, -1}
	}
	left, right := k, k
	for left > 0 && nums[left-1] == target {
		left = left - 1
	}

	for right < len(nums)-1 && nums[right+1] == target {
		right = right + 1
	}
	return []int{left, right}
}
