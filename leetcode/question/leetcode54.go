package question

func SpiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return make([]int, 0)
	}

	var (
		row, columns             = len(matrix), len(matrix[0])
		result                   = make([]int, row*columns)
		index                    = 0
		left, right, top, bottom = 0, columns - 1, 0, row - 1
	)

	for left <= right && top <= bottom {
		for x := left; x <= right; x++ {
			result[index] = matrix[top][x]
			index++
		}
		for y := top + 1; y <= bottom; y++ {
			result[index] = matrix[y][right]
			index++
		}
		if left < right && top < bottom {
			for x := right - 1; x > left; x-- {
				result[index] = matrix[bottom][x]
				index++
			}
			for y := bottom; y > top; y-- {
				result[index] = matrix[y][left]
				index++
			}
		}
		left++
		right--
		top++
		bottom--
	}
	return result
}
