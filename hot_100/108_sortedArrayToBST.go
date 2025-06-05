package main

func sortedArrayToBST(nums []int) *TreeNode {
	return generateTreeNode(nums, 0, len(nums)-1)
}

func generateTreeNode(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right) / 2
	node := &TreeNode{Val: nums[mid]}
	node.Left = generateTreeNode(nums, left, mid-1)
	node.Right = generateTreeNode(nums, mid+1, right)
	return node
}
