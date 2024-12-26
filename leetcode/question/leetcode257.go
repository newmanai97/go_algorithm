package question

import (
	"strconv"
)

func BinaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return make([]string, 0)
	}
	result := make([]string, 0)
	path := make([]int, 0)
	var traversal func(node *TreeNode, path []int)
	traversal = func(node *TreeNode, path []int) {
		path = append(path, node.Val)
		if node.Left == nil && node.Right == nil {
			s := ""
			for i, p := range path {
				s = s + strconv.Itoa(p)
				if i < len(path)-1 {
					s = s + "->"
				}
			}
			result = append(result, s)
			return
		}

		if node.Left != nil {
			traversal(node.Left, path)
		}

		if node.Right != nil {
			traversal(node.Right, path)
		}
	}
	traversal(root, path)
	return result
}
