package question

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	anses := make([][]int, 0)
	anses = append(anses, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= anses[len(anses)-1][1] {
			if intervals[i][1] > anses[len(anses)-1][1] {
				anses[len(anses)-1][1] = intervals[i][1]
			}
		} else {
			anses = append(anses, intervals[i])
		}
	}
	return anses
}
