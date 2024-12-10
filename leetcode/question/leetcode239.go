package question

import "suanfa/leetcode/structure"

func MaxSlidingWindow(nums []int, k int) []int {
	queue := structure.NewQueue()
	len := len(nums)
	result := make([]int, 0)

	for i := 0; i < len; i++ {
		queue.Push(nums[i])
		if i >= k-1 {
			result = append(result, queue.Front())
			queue.Pop(nums[i-k+1])
		}
	}
	return result
}
