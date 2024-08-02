package question

import "sort"

func SortedSquares(nums []int) []int {
	result := make([]int, len(nums))
	k := len(nums) - 1
	for i, j := 0, len(nums)-1; i <= j; {
		if nums[i]*nums[i] > nums[j]*nums[j] {
			result[k] = nums[i] * nums[i]
			k--
			i++
		} else {
			result[k] = nums[j] * nums[j]
			k--
			j--
		}
	}
	return result
}

func SortedSquares2(nums []int) []int {
	for i, j := range nums {
		nums[i] *= j
	}
	sort.Ints(nums)
	return nums
}
