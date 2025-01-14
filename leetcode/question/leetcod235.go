package question

func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val > p.Val && root.Val > q.Val {
		left := lowestCommonAncestor1(root.Left, p, q)
		if left != nil {
			return left
		}
	}
	if root.Val < p.Val && root.Val < q.Val {
		right := lowestCommonAncestor1(root.Right, p, q)
		if right != nil {
			return right
		}
	}
	return root
}

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	for root != nil {
		if root.Val > p.Val && root.Val > q.Val {
			root = root.Left
		} else if root.Val < p.Val && root.Val < q.Val {
			root = root.Right
		} else {
			return root
		}

	}
	return nil
}
