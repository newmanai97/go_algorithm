package question

import (
	"container/list"
	"suanfa/leetcode/structure"
)

func rightSideView(root *structure.TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	stack := list.New()
	stack.PushBack(root)

	for stack.Len() > 0 {
		depth := stack.Len()
		for i := 0; i < depth; i++ {
			node := stack.Remove(stack.Front()).(*structure.TreeNode)
			if node.Left != nil {
				stack.PushBack(node.Left)
			}
			if node.Right != nil {
				stack.PushBack(node.Right)
			}
			if i == depth-1 {
				result = append(result, node.Val)
			}
		}
	}
	return result
}
