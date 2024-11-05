package question

import "strings"

func RepeatedSubstringPattern(s string) bool {
	str := s + s
	str = str[1 : len(str)-1]
	return strings.Contains(str, s)
}

func RepeatedSubstringPattern2(s string) bool {
	next := make([]int, len(s))
	GetNext(next, s)
	maxlen := next[len(next)-1]
	if maxlen == 0 {
		return false
	}
	return len(s)%(len(s)-maxlen) == 0
}
