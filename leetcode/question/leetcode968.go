package question

func MinCameraCover(root *TreeNode) int {
	ans := 0
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 2
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		if left == 2 && right == 2 {
			return 0
		} else if left == 0 || right == 0 {
			ans++
			return 1
		} else {
			return 2
		}
	}

	i := dfs(root)
	if i == 0 {
		ans++
	}
	return ans
}
