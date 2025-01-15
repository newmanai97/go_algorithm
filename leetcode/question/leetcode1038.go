package question

func convertBST2(root *TreeNode) *TreeNode {
	//其实本质上是找到最大值，从大到小累加
	//所以用反向的中序遍历，就是从大到小
	pre := 0
	var convert func(node *TreeNode)
	convert = func(node *TreeNode) {
		if node == nil {
			return
		}
		convert(node.Right)
		node.Val = node.Val + pre
		pre = node.Val
		convert(node.Left)
	}
	convert(root)
	return root
}
