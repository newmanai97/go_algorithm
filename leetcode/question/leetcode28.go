package question

func StrStr(haystack string, needle string) int {
	length := len(needle)
	if length == 0 {
		return -1
	}
	j := 0
	next := make([]int, length)
	GetNext(next, needle)
	for i := 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == length {
			return i - length + 1
		}
	}
	return -1
}

func GetNext(next []int, s string) {
	j := 0
	next[0] = 0

	for i := 1; i < len(s); i++ {
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}
		if s[j] == s[i] {
			j++
		}
		next[i] = j
	}
}
