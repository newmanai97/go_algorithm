package question

import (
	"container/list"
)

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
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
			node := stack.Remove(stack.Front()).(*Node)
			for _, n := range node.Children {
				if n != nil {
					stack.PushBack(n)
				}
			}
			res = append(res, node.Val)
		}
		result = append(result, res)
	}
	return result
}
