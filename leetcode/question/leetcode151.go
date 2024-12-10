package question

import (
	"strings"
)

func ReverseWords(s string) string {
	s = strings.TrimSpace(s)

	bytes := []byte(s)

	result := []byte{}
	for i, j := len(bytes)-1, len(bytes)-1; i >= 0; i-- {
		if bytes[i] == ' ' {
			if j == len(bytes)-1 {
				result = append(result, bytes[i+1:]...)
			} else {
				result = append(result, bytes[i+1:j+1]...)
			}
			result = append(result, ' ')

			j = i
			for ; j >= 0; j-- {
				if bytes[j] != ' ' {
					break
				}
			}
			i = j + 1
		}
		if i == 0 {
			result = append(result, bytes[:j+1]...)
		}
	}

	return string(result)
}

func ReverseWords2(s string) string {
	slow := 0
	bytes := []byte(s)

	for i := 0; i < len(bytes); i++ {
		if bytes[i] != ' ' {
			if slow != 0 {
				bytes[slow] = ' '
				slow++
			}
			for i < len(bytes) && bytes[i] != ' ' {
				bytes[slow] = bytes[i]
				slow++
				i++
			}
		}
	}

	bytes = bytes[:slow]

	reverse(bytes, 0, len(bytes)-1)
	last := 0
	for i := 0; i <= len(bytes); i++ {
		if i == len(bytes) || bytes[i] == ' ' {
			reverse(bytes, last, i-1)
			last = i + 1
		}
	}
	return string(bytes)
}
