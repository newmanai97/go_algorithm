package question

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	ans := 1
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= intervals[i-1][1] {
			ans++
		} else {
			if intervals[i-1][1] < intervals[i][1] {
				intervals[i][1] = intervals[i-1][1]
			}
		}
	}
	return len(intervals) - ans
}
