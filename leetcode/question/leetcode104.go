package question

import (
	"container/list"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 层序遍历法
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	stack := list.New()
	stack.PushBack(root)

	var result = 0

	for stack.Len() > 0 {
		depth := stack.Len()
		for i := 0; i < depth; i++ {
			node := stack.Remove(stack.Front()).(*TreeNode)
			if node.Left != nil {
				stack.PushBack(node.Left)
			}
			if node.Right != nil {
				stack.PushBack(node.Right)
			}
		}
		result++
	}
	return result
}

// 递归法
func maxDepth2(root *TreeNode) int {
	var getMaxDepth func(node *TreeNode) int
	getMaxDepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		depthLeft := getMaxDepth(node.Left)
		depthRight := getMaxDepth(node.Right)
		depth := max(depthLeft, depthRight) + 1
		return depth
	}
	return getMaxDepth(root)
}
