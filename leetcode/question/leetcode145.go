package question

import "suanfa/leetcode/structure"

// 递归法
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

// 迭代法
func PostorderTraversal2(root *structure.TreeNode) []int {
	var result []int

	stack := make([]*structure.TreeNode, 0)
	if root == nil {
		return result
	}
	stack = append(stack, root)

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		result = append(result, node.Val)
		stack = stack[:len(stack)-1]

		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	left := 0
	right := len(result) - 1

	for left < right {
		result[left], result[right] = result[right], result[left]
		left++
		right--
	}

	return result
}

// 统一迭代法
func PostorderTraversal3(root *structure.TreeNode) []int {
	var result []int

	stack := make([]*structure.TreeNode, 0)
	if root == nil {
		return result
	}
	stack = append(stack, root)

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node != nil {
			stack = append(stack, node)
			stack = append(stack, nil)
			if node.Right != nil {
				stack = append(stack, node.Right)
			}

			if node.Left != nil {
				stack = append(stack, node.Left)
			}

		} else {
			res := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = append(result, res.Val)
		}
	}
	return result
}
