package question

import "container/list"

// 递归法
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	var merge func(t1 *TreeNode, t2 *TreeNode) *TreeNode
	merge = func(t1 *TreeNode, t2 *TreeNode) *TreeNode {
		if t1 == nil {
			return t2
		}

		if t2 == nil {
			return t1
		}

		t1.Val = t1.Val + t2.Val
		t1.Left = merge(t1.Left, t2.Left)
		t1.Right = merge(t1.Right, t2.Right)
		return t1
	}
	return merge(root1, root2)
}

// 迭代法
func mergeTrees2(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	queue := list.New()
	if root1 == nil {
		return root2
	}

	if root2 == nil {
		return root1
	}

	queue.PushBack(root1)
	queue.PushBack(root2)

	for size := queue.Len(); size > 0; size = queue.Len() {
		node1 := queue.Remove(queue.Front()).(*TreeNode)
		node2 := queue.Remove(queue.Front()).(*TreeNode)

		node1.Val = node1.Val + node2.Val

		if node1.Left != nil && node2.Left != nil {
			queue.PushBack(node1.Left)
			queue.PushBack(node2.Left)
		}

		if node1.Right != nil && node2.Right != nil {
			queue.PushBack(node1.Right)
			queue.PushBack(node2.Right)
		}
		if node1.Left == nil {
			node1.Left = node2.Left
		}

		if node1.Right == nil {
			node1.Right = node2.Right
		}
	}
	return root1
}
