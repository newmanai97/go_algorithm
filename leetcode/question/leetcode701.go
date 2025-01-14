package question

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val: val,
		}
	}

	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	}

	if root.Val < val {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}

func insertIntoBST2(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val: val,
		}
	}
	node := root
	var parent *TreeNode
	for node != nil {
		parent = node
		if node.Val > val {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	if parent.Val > val {
		parent.Left = &TreeNode{
			Val: val,
		}
	} else {
		parent.Right = &TreeNode{
			Val: val,
		}
	}
	return root
}
