package question

import "strconv"

func monotoneIncreasingDigits(n int) int {
	str := strconv.Itoa(n)
	for i := len(str) - 2; i >= 0; i-- {
		if str[i] > str[i+1] {
			str = str[:i] + string(str[i]-1) + str[i+1:]
			for j := i + 1; j < len(str); j++ {
				str = str[:j] + "9" + str[j+1:]
			}
		}
	}
	ans, _ := strconv.Atoi(str)
	return ans
}
