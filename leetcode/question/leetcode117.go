package question

import (
	"suanfa/leetcode/structure"
)

func Connect2(root *structure.Node2) *structure.Node2 {
	if root == nil {
		return root
	}
	queue := make([]*structure.Node2, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[i]
			if i != size-1 {
				node.Next = queue[i+1]
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[size:]
	}
	return root
}
