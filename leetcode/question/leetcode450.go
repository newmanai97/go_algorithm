package question

func deleteNode(root *TreeNode, key int) *TreeNode {

	var delete func(node *TreeNode, key int) *TreeNode
	delete = func(node *TreeNode, key int) *TreeNode {
		//终止条件
		if node == nil {
			return nil
		}
		if node.Val == key {
			if node.Left == nil && node.Right == nil {
				return nil
			} else if node.Left != nil && node.Right == nil {
				return node.Left
			} else if node.Left == nil && node.Right != nil {
				return node.Right
			} else {
				cur := node.Right
				for cur.Left != nil {
					cur = cur.Left
				}
				cur.Left = node.Left
				return node.Right
			}
		}

		//递归流程
		if node.Val > key {
			node.Left = delete(node.Left, key)
		}
		if node.Val < key {
			node.Right = delete(node.Right, key)
		}
		return node
	}
	return delete(root, key)
}
