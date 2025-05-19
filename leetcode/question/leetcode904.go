package question

func totalFruit(fruits []int) int {
	tmp := map[int]int{}
	result := 0
	i := 0
	for j, x := range fruits {
		tmp[x]++
		for len(tmp) > 2 {
			y := fruits[i]
			tmp[y]--
			if tmp[y] == 0 {
				delete(tmp, y)
			}
			i++
		}
		if j-i+1 > result {
			result = j - i + 1
		}
	}
	return result
}
