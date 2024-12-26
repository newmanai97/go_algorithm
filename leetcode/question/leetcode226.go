package question

func invertTree(root *TreeNode) *TreeNode {

	var invert func(node *TreeNode)
	invert = func(node *TreeNode) {
		if node == nil {
			return
		}
		node.Left, node.Right = node.Right, node.Left
		invert(node.Left)
		invert(node.Right)
	}
	invert(root)
	return root
}
