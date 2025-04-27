package question

// 二分查找
// 左闭右闭
func search(nums []int, target int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		mid := (i + j) / 2
		if nums[mid] == target {
			return mid
		} else {
			if nums[mid] > target {
				j = mid - 1
			} else {
				i = mid + 1
			}
		}
	}
	return -1
}

// 左闭右开
func search2(nums []int, target int) int {
	i, j := 0, len(nums)
	for i < j {
		mid := (i + j) / 2
		if nums[mid] == target {
			return mid
		} else {
			if nums[mid] > target {
				j = mid
			} else {
				i = mid + 1
			}
		}
	}
	return -1
}
