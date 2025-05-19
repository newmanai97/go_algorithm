package question

func generateMatrix(n int) [][]int {
	loop := n / 2
	startx, starty, offset := 0, 0, 1
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}
	mid := n / 2
	count := 1
	i, j := 0, 0
	for loop > 0 {
		i, j = startx, starty

		for j < n-offset {
			result[i][j] = count
			count++
			j++
		}

		for i < n-offset {
			result[i][j] = count
			count++
			i++
		}
		for j > starty {
			result[i][j] = count
			count++
			j--
		}
		for i > startx {
			result[i][j] = count
			count++
			i--
		}
		startx++
		starty++
		offset++
		loop--
	}
	if n%2 != 0 {
		result[mid][mid] = count
	}
	return result
}
