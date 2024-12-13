package question

import "suanfa/leetcode/structure"

func InorderTraversal(root *structure.TreeNode) []int {
	var result []int
	var traversal func(node *structure.TreeNode)
	traversal = func(node *structure.TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)
		result = append(result, node.Val)
		traversal(node.Right)
	}
	traversal(root)
	return result
}
