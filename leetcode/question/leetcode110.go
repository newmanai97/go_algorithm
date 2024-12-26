package question

import "math"

func isBalanced(root *TreeNode) bool {

	var getheight func(node *TreeNode) int
	getheight = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftheight := getheight(node.Left)
		if leftheight == -1 {
			return -1
		}

		rightheight := getheight(node.Right)
		if rightheight == -1 {
			return -1
		}

		if math.Abs(float64(rightheight-leftheight)) > 1 {
			return -1
		}

		return int(math.Max(float64(leftheight), float64(rightheight))) + 1
	}
	result := getheight(root) == -1
	if result {
		return false
	}
	return true
}
