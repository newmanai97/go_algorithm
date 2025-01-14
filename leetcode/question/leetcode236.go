package question

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	var traverse func(node, p, q *TreeNode) *TreeNode
	traverse = func(node, p, q *TreeNode) *TreeNode {
		if node == nil || node == p || node == q {
			return node
		}
		left := traverse(node.Left, p, q)
		right := traverse(node.Right, p, q)
		if left != nil && right != nil {
			return node
		} else if left != nil && right == nil {
			return left
		} else if left == nil && right != nil {
			return right
		} else {
			return nil
		}
	}
	return traverse(root, p, q)
}
