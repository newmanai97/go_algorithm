package question

import "suanfa/leetcode/structure"

// 递归法
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

// 迭代法
func InorderTraversal2(root *structure.TreeNode) []int {
	var result []int

	if root == nil {
		return result
	}
	stack := make([]*structure.TreeNode, 0)
	cur := root
	for cur != nil || len(stack) != 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			cur = stack[len(stack)-1]
			result = append(result, cur.Val)
			stack = stack[:len(stack)-1]
			cur = cur.Right
		}
	}
	return result
}
