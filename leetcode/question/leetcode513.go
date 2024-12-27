package question

import (
	"container/list"
	"math"
)

// 层序遍历
func FindBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := list.New()
	queue.PushBack(root)

	leftvalue := 0

	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			//从队列头弹出元素
			node := queue.Remove(queue.Front()).(*TreeNode)
			//获取第一个元素的值
			if i == 0 {
				leftvalue = node.Val
			}
			//将该节点的左右孩子入队列
			if node.Left != nil {
				queue.PushBack(node.Left)
			}

			if node.Right != nil {
				queue.PushBack(node.Right)
			}

		}
	}
	return leftvalue
}

// 递归法
func FindBottomLeftValue2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxdepth, result := math.MinInt, 0
	var getleftvalue func(node *TreeNode, depth int)
	getleftvalue = func(node *TreeNode, depth int) {
		if node != nil && node.Left == nil && node.Right == nil {
			if depth > maxdepth {
				maxdepth = depth
				result = node.Val
			}
		}

		if node.Left != nil {
			depth++
			getleftvalue(node.Left, depth)
			depth--
		}

		if node.Right != nil {
			depth++
			getleftvalue(node.Right, depth)
			depth--
		}
		return
	}
	getleftvalue(root, 0)
	return result
}
