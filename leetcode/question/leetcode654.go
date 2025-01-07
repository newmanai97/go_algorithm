package question

func constructMaximumBinaryTree(nums []int) *TreeNode {

	var getMaximumBinaryTree func(param []int) *TreeNode
	getMaximumBinaryTree = func(param []int) *TreeNode {
		//如果数组只有一个节点，那就直接返回
		if len(param) == 1 {
			return &TreeNode{
				Val: param[0],
			}
		}

		//获取数组内的最大值
		index := 0
		maxvalue := param[index]
		for i := 0; i < len(param); i++ {
			if param[i] > maxvalue {
				index = i
				maxvalue = param[i]
			}
		}
		//构造根节点
		root := &TreeNode{
			Val: maxvalue,
		}

		//切割左右子树数组
		left := param[:index]
		right := param[index+1:]

		//递归
		if len(left) != 0 {
			root.Left = getMaximumBinaryTree(left)
		}
		if len(right) != 0 {
			root.Right = getMaximumBinaryTree(right)
		}
		return root
	}
	return getMaximumBinaryTree(nums)
}
