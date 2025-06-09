package question

import "sort"

func findMinArrowShots(points [][]int) int {

	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	ans := 1
	for i := 1; i < len(points); i++ {
		if points[i-1][1] < points[i][0] {
			ans++
		} else {
			if points[i-1][1] < points[i][1] {
				points[i][1] = points[i-1][1]
			}
		}
	}

	return ans
}
