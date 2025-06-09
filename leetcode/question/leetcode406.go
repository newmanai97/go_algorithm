package question

import "sort"

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})

	for i, v := range people {
		copy(people[v[1]+1:i+1], people[v[1]:i+1])
		people[v[i]] = v
	}
	return people
}
