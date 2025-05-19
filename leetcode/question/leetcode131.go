package question

var (
	path6 []string
	res6  [][]string
)

func partition(s string) [][]string {
	path6, res6 = make([]string, 0, len(s)), make([][]string, 0)
	dfs6(s, 0)
	return res6
}

func dfs6(s string, start int) {
	if start == len(s) {
		if len(path6) != 0 {
			tmp := make([]string, len(path6))
			copy(tmp, path6)
			res6 = append(res6, tmp)
			return
		}
	}
	for i := start; i < len(s); i++ {
		str := s[start : i+1]
		if isPalindrome(str) {
			path6 = append(path6, str)
			dfs6(s, i+1)
			path6 = path6[:len(path6)-1]
		}
	}
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i <= j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
