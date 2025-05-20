package question

import (
	"strconv"
	"strings"
)

var (
	path7 []string
	res7  []string
)

func restoreIpAddresses(s string) []string {
	path7, res7 = make([]string, 0, 4), make([]string, 0)
	dfs7(s, 0)
	return res7
}

func dfs7(s string, start int) {
	if len(path7) == 4 {
		if start >= len(s) {
			str := strings.Join(path7, ".")
			res7 = append(res7, str)
		}
		return
	}

	for i := start; i < len(s); i++ {
		if i != start && s[start] == '0' {
			break
		}
		str := s[start : i+1]
		num, _ := strconv.Atoi(str)
		if num >= 0 && num <= 255 {
			path7 = append(path7, strconv.Itoa(num))
			dfs7(s, i+1)
			path7 = path7[:len(path7)-1]
		} else {
			break
		}
	}
}
