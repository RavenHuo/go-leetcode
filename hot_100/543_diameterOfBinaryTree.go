package main

// 左右子树之和
func diameterOfBinaryTree(root *TreeNode) int {
	res := 0
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return -1
		}
		// left := dfs(node.Left)
		// right := dfs(node.Right)
		// // 以当前节点拐弯(做连接)的链长 = 左子树的链长 + 右子树的链长 + 2(2是该节点连接左右子树需要的长度)
		// // 对于叶子节点，left和right都是-1，那么它的链长为0，符合预期
		// res = max(res, left+right+2)
		// // 最后的返回值，+1是因为，无论选择左子树还是右子树，最后到当前节点都还有1的距离
		// return max(left, right)+1

		// 也可以写成下面这样（左右子树链长自+1来找当前父节点）
		left := dfs(node.Left) + 1
		right := dfs(node.Right) + 1
		res = max(res, left+right)
		return max(left, right)
	}
	dfs(root)
	return res
}
