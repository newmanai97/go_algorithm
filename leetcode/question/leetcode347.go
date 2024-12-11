package question

import "sort"

func TopKFrequent(nums []int, k int) []int {
	result := make([]int, 0)
	resultMap := map[int]int{}

	for _, j := range nums {
		resultMap[j]++
	}

	for key, _ := range resultMap {
		result = append(result, key)
	}

	sort.Slice(result, func(a, b int) bool {
		return resultMap[result[a]] > resultMap[result[b]]
	})

	return result[:k]
}
