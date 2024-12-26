package question

func SumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var getleftleavesval func(node *TreeNode) int
	getleftleavesval = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftval := getleftleavesval(node.Left)
		if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
			leftval = node.Left.Val
		}
		rightval := getleftleavesval(node.Right)
		return leftval + rightval
	}
	return getleftleavesval(root)
}
