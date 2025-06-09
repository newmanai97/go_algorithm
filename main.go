package main

import (
	"suanfa/leetcode/question"
)

func main() {
	node1 := &question.TreeNode{Val: 0}
	node2 := &question.TreeNode{Val: 0}
	node3 := &question.TreeNode{Val: 0}
	node4 := &question.TreeNode{Val: 0}
	node5 := &question.TreeNode{Val: 0}
	node1.Left = node2
	node2.Left = node3
	node3.Left = node4
	node4.Right = node5
	i := question.MinCameraCover(node1)

	println("i is %v", i)
}
