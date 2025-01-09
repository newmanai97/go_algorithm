package question

func isValidBST(root *TreeNode) bool {

	var pre *TreeNode
	var isValid func(node *TreeNode) bool
	isValid = func(node *TreeNode) bool {
		if node == nil {
			return true
		}

		left := isValid(node.Left)
		if pre != nil && node.Val <= pre.Val {
			return false
		}
		pre = node

		right := isValid(node.Right)

		return left && right
	}
	return isValid(root)
}

func isValidBST2(root *TreeNode) bool {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return true
	}

	var arr []int
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)
		arr = append(arr, node.Val)
		traversal(node.Right)
	}
	traversal(root)

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] >= arr[i+1] {
			return false
		}
	}
	return true
}
