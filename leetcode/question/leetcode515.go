package question

import (
	"container/list"
	"math"
	"suanfa/leetcode/structure"
)

func largestValues(root *structure.TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	stack := list.New()
	stack.PushBack(root)

	for stack.Len() > 0 {
		max := math.MinInt
		depth := stack.Len()
		for i := 0; i < depth; i++ {
			node := stack.Remove(stack.Front()).(*structure.TreeNode)
			if node.Left != nil {
				stack.PushBack(node.Left)
			}
			if node.Right != nil {
				stack.PushBack(node.Right)
			}
			if max < node.Val {
				max = node.Val
			}
		}
		result = append(result, max)
	}
	return result
}
