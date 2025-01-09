package question

func findMode(root *TreeNode) []int {
	value, maxvalue := 0, 0
	result := make([]int, 0)
	var prev *TreeNode

	var find func(node *TreeNode)
	find = func(node *TreeNode) {
		if node == nil {
			return
		}
		find(node.Left)

		if prev == nil {
			value = 1
		} else if prev != nil && prev.Val == node.Val {
			value++
		} else {
			value = 1
		}

		prev = node

		if value == maxvalue {
			result = append(result, node.Val)
		}

		if value > maxvalue {
			maxvalue = value
			result = make([]int, 0)
			result = append(result, node.Val)
		}

		find(node.Right)
	}
	find(root)
	return result
}
