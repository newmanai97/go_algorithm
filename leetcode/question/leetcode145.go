package question

import "suanfa/leetcode/structure"

func PostorderTraversal(root *structure.TreeNode) []int {
	var result []int
	var traversal func(node *structure.TreeNode)
	traversal = func(node *structure.TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)
		traversal(node.Right)
		result = append(result, node.Val)
	}
	traversal(root)
	return result
}
