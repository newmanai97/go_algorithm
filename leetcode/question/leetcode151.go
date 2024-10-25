package question

import "strings"

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
