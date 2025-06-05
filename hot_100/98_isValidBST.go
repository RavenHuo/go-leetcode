package main

// 上下界
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isValidBstWithBound(root, nil, nil)
}
func isValidBstWithBound(root *TreeNode, min, max *int) bool {
	if root == nil {
		return true
	}
	if max != nil && root.Val >= *max {
		return false
	}
	if min != nil && root.Val <= *min {
		return false
	}

	return isValidBstWithBound(root.Left, min, &root.Val) && isValidBstWithBound(root.Right, &root.Val, max)
}
