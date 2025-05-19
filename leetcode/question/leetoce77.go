package question

var (
	path1  []int
	result [][]int
)

func combine(n int, k int) [][]int {
	path1, result = make([]int, 0, k), make([][]int, 0)
	dfs1(n, k, 1)
	return result
}

func dfs1(n, k, startIndex int) {
	if len(path1) == k {
		tmp := make([]int, k)
		copy(tmp, path1)
		result = append(result, tmp)
		return
	}
	for i := startIndex; i <= n-(k-len(path1))+1; i++ {
		path1 = append(path1, i)
		dfs1(n, k, i+1)
		path1 = path1[:len(path1)-1]
	}

}
