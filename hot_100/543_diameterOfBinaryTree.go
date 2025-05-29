package main

// 左右子树之和
func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0
	getDepthAndUpdateDiameter(root, &maxDiameter)
	return maxDiameter
}

// 抽出的递归函数
func getDepthAndUpdateDiameter(node *TreeNode, maxDiameter *int) int {
	if node == nil {
		return 0
	}
	left := getDepthAndUpdateDiameter(node.Left, maxDiameter)
	right := getDepthAndUpdateDiameter(node.Right, maxDiameter)

	// 更新直径：左右子树深度之和
	if left+right > *maxDiameter {
		*maxDiameter = left + right
	}

	// 返回当前子树的最大深度
	return 1 + max(left, right)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
