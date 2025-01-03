package question

func pathSum(root *TreeNode, targetSum int) [][]int {
	result := make([][]int, 0)

	if root == nil {
		return result
	}

	var getpath func(node *TreeNode, target int, path []int)
	getpath = func(node *TreeNode, target int, path []int) {
		if node.Left == nil && node.Right == nil && target == 0 {
			copypath := make([]int, 0)
			for _, n := range path {
				copypath = append(copypath, n)
			}
			result = append(result, copypath)
		}
		if node.Right == nil && node.Left == nil {
			return
		}
		if node.Left != nil {
			path = append(path, node.Left.Val)
			target = target - node.Left.Val
			getpath(node.Left, target, path)
			target = target + node.Left.Val
			path = path[:len(path)-1]
		}

		if node.Right != nil {
			path = append(path, node.Right.Val)
			target = target - node.Right.Val
			getpath(node.Right, target, path)
			target = target + node.Right.Val
			path = path[:len(path)-1]
		}
	}
	path := make([]int, 0)
	path = append(path, root.Val)
	getpath(root, targetSum-root.Val, path)
	return result
}
