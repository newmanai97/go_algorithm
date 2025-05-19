package question

var (
	m     []string
	path3 []byte
	res3  []string
)

func letterCombinations(digits string) []string {
	m = []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	path3, res3 = make([]byte, 0), make([]string, 0)
	if len(digits) == 0 {
		return res3
	}
	dfs3(digits, 0)
	return res3
}

func dfs3(digits string, start int) {
	if len(path3) == len(digits) {
		tmp := make([]byte, 0)
		copy(tmp, path3)
		res3 = append(res3, string(path3))
		return
	}
	str := m[int(digits[start]-'0')-2]
	for i := 0; i < len(str); i++ {
		path3 = append(path3, str[i])
		dfs3(digits, start+1)
		path3 = path3[:len(path3)-1]
	}
}
