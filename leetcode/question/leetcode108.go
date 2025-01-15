package question

func sortedArrayToBST(nums []int) *TreeNode {
	var sorted func(nums []int, i, j int) *TreeNode
	sorted = func(nums []int, i, j int) *TreeNode {
		if i > j {
			return nil
		}

		mid := (i + j) / 2
		midNode := &TreeNode{
			Val: nums[mid],
		}
		midNode.Left = sorted(nums, i, mid-1)
		midNode.Right = sorted(nums, mid+1, j)
		return midNode
	}
	return sorted(nums, 0, len(nums)-1)
}
