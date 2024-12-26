package question

import "container/list"

func minDepth(root *TreeNode) int {
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
			if node.Left == nil && node.Right == nil {
				result++
				return result
			}
		}
		result++
	}
	return result
}
