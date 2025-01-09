package question

import "math"

func getMinimumDifference(root *TreeNode) int {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return 0
	}

	var arr []int
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)
		arr = append(arr, node.Val)
		traversal(node.Right)
	}
	traversal(root)

	minvalue := math.MaxInt
	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] < minvalue {
			minvalue = arr[i+1] - arr[i]
		}
	}
	return minvalue
}
