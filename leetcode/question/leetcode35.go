package question

func searchInsert(nums []int, target int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		mid := (i + j) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			j = mid - 1
			if j < 0 || nums[j] < target {
				return mid
			}
		} else {
			i = mid + 1
			if i > len(nums)-1 || nums[i] > target {
				return i
			}
		}
	}
	return -1
}
