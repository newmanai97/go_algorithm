package question

import (
	"container/list"
	"suanfa/leetcode/structure"
)

func averageOfLevels(root *structure.TreeNode) []float64 {
	result := make([]float64, 0)
	if root == nil {
		return result
	}
	stack := list.New()
	stack.PushBack(root)

	for stack.Len() > 0 {
		res := 0
		depth := stack.Len()
		for i := 0; i < depth; i++ {
			node := stack.Remove(stack.Front()).(*structure.TreeNode)
			if node.Left != nil {
				stack.PushBack(node.Left)
			}
			if node.Right != nil {
				stack.PushBack(node.Right)
			}
			res = res + node.Val
		}
		result = append(result, float64(res)/float64(depth))
	}
	return result
}
