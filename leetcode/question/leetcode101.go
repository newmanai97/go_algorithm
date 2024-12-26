package question

import "container/list"

func isSymmetric(root *TreeNode) bool {
	var compare func(left, right *TreeNode) bool
	compare = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		} else if left == nil || right == nil {
			return false
		} else if left.Val != right.Val {
			return false
		}
		return compare(left.Left, left.Right) && compare(right.Left, right.Right)
	}
	return compare(root.Left, root.Right)
}

// 迭代法
func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := list.New()
	queue.PushBack(root.Left)
	queue.PushBack(root.Right)

	for queue.Len() > 0 {
		left := queue.Remove(queue.Front()).(*TreeNode)
		right := queue.Remove(queue.Front()).(*TreeNode)
		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}
		queue.PushBack(left.Left)
		queue.PushBack(right.Right)
		queue.PushBack(left.Right)
		queue.PushBack(right.Left)
	}
	return true
}
