package question

import "suanfa/leetcode/structure"

func PreorderTraversal(root *structure.TreeNode) []int {
	var result []int
	var traversal func(node *structure.TreeNode)
	traversal = func(node *structure.TreeNode) {
		if node == nil {
			return
		}
		result = append(result, node.Val)
		traversal(node.Left)
		traversal(node.Right)
	}
	traversal(root)
	return result

}
