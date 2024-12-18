package question

import "suanfa/leetcode/structure"

// 递归方法
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

// 迭代方法
func PreorderTraversal2(root *structure.TreeNode) []int {
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

		if node.Right != nil {
			stack = append(stack, node.Right)
		}

		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return result
}

// 统一迭代法
func PreorderTraversal3(root *structure.TreeNode) []int {
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
			if node.Right != nil {
				stack = append(stack, node.Right)
			}

			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			stack = append(stack, node)
			stack = append(stack, nil)

		} else {
			res := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = append(result, res.Val)
		}
	}
	return result
}
