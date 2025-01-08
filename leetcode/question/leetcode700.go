package question

// 递归法
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	var result *TreeNode
	if root.Val == val {
		result = root
	} else if root.Val > val {
		result = searchBST(root.Left, val)
	} else {
		result = searchBST(root.Right, val)
	}
	return result
}

// 迭代法
func searchBST2(root *TreeNode, val int) *TreeNode {
	for root != nil {
		if root.Val == val {
			return root
		} else if root.Val > val {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return nil
}
