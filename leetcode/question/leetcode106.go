package question

func BuildTree(inorder []int, postorder []int) *TreeNode {
	if postorder == nil || len(postorder) == 0 {
		return nil
	}

	var method func(in []int, post []int) *TreeNode
	method = func(in []int, post []int) *TreeNode {
		//递归结束条件
		if post == nil || len(post) == 0 {
			return nil
		}

		//后续遍历的最后一个值就是根节点
		rootvalue := post[len(post)-1]
		root := &TreeNode{
			Val: rootvalue,
		}
		post = post[:len(post)-1]
		//如果根节点是叶子节点就直接返回
		if len(post) == 0 {
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
		leftpost := post[:len(leftin)]
		rightpost := post[len(leftin):]

		//执行递归
		root.Left = method(leftin, leftpost)
		root.Right = method(rightin, rightpost)
		return root
	}

	return method(inorder, postorder)
}
