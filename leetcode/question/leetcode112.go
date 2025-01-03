package question

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	var judgepath func(node *TreeNode, target int) bool
	judgepath = func(node *TreeNode, target int) bool {
		if node.Left == nil && node.Right == nil && target == 0 {
			return true
		}
		if node.Left == nil && node.Right == nil && target != 0 {
			return false
		}
		if node.Left != nil {
			target = target - node.Left.Val
			if judgepath(node.Left, target) {
				return true
			}
			target = target + node.Left.Val
		}

		if node.Right != nil {
			target = target - node.Right.Val
			if judgepath(node.Right, target) {
				return true
			}
			target = target + node.Right.Val
		}
		return false
	}
	return judgepath(root, targetSum-root.Val)
}
