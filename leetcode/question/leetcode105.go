package question

func buildTree(preorder []int, inorder []int) *TreeNode {
	if preorder == nil || len(preorder) == 0 {
		return nil
	}

	var method func(pre []int, in []int) *TreeNode
	method = func(pre []int, in []int) *TreeNode {
		//递归结束条件
		if pre == nil || len(pre) == 0 {
			return nil
		}

		//后续遍历的最后一个值就是根节点
		rootvalue := pre[0]
		root := &TreeNode{
			Val: rootvalue,
		}
		pre = pre[1:]
		//如果根节点是叶子节点就直接返回
		if len(pre) == 0 {
			return root
		}

		//根据根节点切割中序数组
		index := 0
		for k, i := range in {
			if i == rootvalue {
				index = k
			}
		}
		leftin := in[:index]
		rightin := in[index+1:]

		//根据中序遍历左子树大小切分后序数组
		leftpre := pre[:len(leftin)]
		rightpre := pre[len(leftin):]

		//执行递归
		root.Left = method(leftpre, leftin)
		root.Right = method(rightpre, rightin)
		return root
	}

	return method(preorder, inorder)
}
