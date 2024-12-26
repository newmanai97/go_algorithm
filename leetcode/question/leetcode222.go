package question

import (
	"container/list"
)

// 递归法--后序遍历
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var count func(node *TreeNode) int
	count = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftcount, rightcount := 0, 0
		if node.Left != nil {
			leftcount = count(node.Left)
		}

		if node.Right != nil {
			rightcount = count(node.Right)
		}

		result := leftcount + rightcount + 1
		return result
	}
	return count(root)
}

// 迭代法--前序遍历
func countNodes2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var result []int

	stack := make([]*TreeNode, 0)
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
	return len(result)
}

// 迭代法--层序遍历
func countNodes3(root *TreeNode) int {
	if root == nil {
		return 0
	}
	result := 0
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		for length > 0 {
			node := queue.Remove(queue.Front()).(*TreeNode)
			length--
			result++
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}
	return result
}

// 递归法 -- 按照满二叉树的性质
func CountNodes5(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var count func(node *TreeNode) int
	count = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left, right := node.Left, node.Right
		leftdepth, rightdepth := 0, 0
		for left != nil {
			left = left.Left
			leftdepth++
		}
		for right != nil {
			right = right.Right
			rightdepth++
		}

		if leftdepth == rightdepth {
			return (2 << leftdepth) - 1
		}

		leftcount := count(node.Left)
		rightcount := count(node.Right)
		result := leftcount + rightcount + 1
		return result
	}

	return count(root)
}
