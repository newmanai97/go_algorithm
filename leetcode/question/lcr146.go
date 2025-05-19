package question

func spiralArray(array [][]int) []int {
	if len(array) == 0 || len(array[0]) == 0 {
		return []int{}
	}
	var (
		row, columns             = len(array), len(array[0])
		result                   = make([]int, row*columns)
		index                    = 0
		left, right, top, bottom = 0, columns - 1, 0, row - 1
	)

	for left <= right && top <= bottom {
		for x := left; x <= right; x++ {
			result[index] = array[top][x]
			index++
		}
		for y := top + 1; y <= bottom; y++ {
			result[index] = array[y][right]
			index++
		}
		if left < right && top < bottom {
			for x := right - 1; x > left; x-- {
				result[index] = array[bottom][x]
				index++
			}
			for y := bottom; y > top; y-- {
				result[index] = array[y][left]
				index++
			}
		}
		left++
		top++
		right--
		bottom--
	}
	return result
}
