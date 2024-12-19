package question

import (
	"container/list"
	"suanfa/leetcode/structure"
)

func levelOrderBottom(root *structure.TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}
	stack := list.New()
	stack.PushBack(root)

	for stack.Len() > 0 {
		res := []int{}
		depth := stack.Len()
		for i := 0; i < depth; i++ {
			node := stack.Remove(stack.Front()).(*structure.TreeNode)
			if node.Left != nil {
				stack.PushBack(node.Left)
			}
			if node.Right != nil {
				stack.PushBack(node.Right)
			}
			res = append(res, node.Val)
		}
		result = append(result, res)
	}

	left, right := 0, len(result)-1
	for left < right {
		result[left], result[right] = result[right], result[left]
	}
	return result
}
