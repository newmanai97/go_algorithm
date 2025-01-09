package question

import "math"

func minDiffInBST(root *TreeNode) int {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return 0
	}

	minvalue := math.MaxInt
	var pre *TreeNode
	var minDiff func(node *TreeNode) int
	minDiff = func(node *TreeNode) int {
		if node == nil {
			return math.MaxInt
		}
		minleft := minDiff(node.Left)
		if pre != nil && node.Val-pre.Val < minvalue {
			minvalue = node.Val - pre.Val
		}
		pre = node
		minright := minDiff(node.Right)
		if minleft < minvalue {
			minvalue = minleft
		}

		if minright < minvalue {
			minvalue = minright
		}
		return minvalue
	}
	return minDiff(root)
}
