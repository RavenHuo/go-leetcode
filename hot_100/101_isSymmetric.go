package main

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}
	return SymmetricTree(root.Left, root.Right)
}

func SymmetricTree(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil && right != nil {
		return false
	}
	if left != nil && right == nil {
		return false
	}
	return left.Val == right.Val && SymmetricTree(left.Right, right.Left) && SymmetricTree(left.Left, right.Right)
}
