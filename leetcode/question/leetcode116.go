package question

import (
	"container/list"
	"suanfa/leetcode/structure"
)

func Connect(root *structure.Node2) *structure.Node2 {
	if root == nil {
		return root
	}
	stack := list.New()
	stack.PushBack(root)

	for stack.Len() > 0 {
		var pre *structure.Node2
		depth := stack.Len()
		for i := 0; i < depth; i++ {
			node := stack.Remove(stack.Front()).(*structure.Node2)
			if node.Left != nil {
				stack.PushBack(node.Left)
			}
			if node.Right != nil {
				stack.PushBack(node.Right)
			}
			if pre == nil {
				pre = node
				continue
			}
			pre.Next = node
			pre = node
		}

	}
	return root
}
