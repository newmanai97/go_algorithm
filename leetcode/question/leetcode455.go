package question

import "sort"

func findContentChildren(g []int, s []int) int {
	num := 0
	sort.Ints(g)
	sort.Ints(s)
	for i := len(g) - 1; i >= 0; i-- {
		for j := 0; j < len(s); j++ {
			if s[j] >= g[i] {
				num++
				s[j] = 0
				break
			}
		}
	}
	return num
}

func findContentChildren2(g []int, s []int) int {
	num := 0
	index := len(s) - 1
	sort.Ints(g)
	sort.Ints(s)
	for i := len(g) - 1; i >= 0; i-- {
		if index >= 0 && s[index] >= g[i] {
			num++
			index--
		}
	}
	return num
}

func findContentChildren3(g []int, s []int) int {
	index := 0
	sort.Ints(g)
	sort.Ints(s)
	for i := 0; i < len(s); i++ {
		if index < len(g) && s[i] >= g[index] {
			index++
		}
	}
	return index
}
