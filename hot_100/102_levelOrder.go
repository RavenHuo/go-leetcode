package main

var levelOrderRes [][]int

func levelOrder(root *TreeNode) [][]int {
	levelOrderRes = make([][]int, 0)
	appendTreeNode(root, 0)
	return levelOrderRes
}
func appendTreeNode(root *TreeNode, index int) {
	if root != nil {
		if len(levelOrderRes) >= index+1 {
			levelOrderRes[index] = append(levelOrderRes[index], root.Val)
		} else {
			levelOrderRes = append(levelOrderRes, make([]int, 0))
			levelOrderRes[index] = append(levelOrderRes[index], root.Val)
		}
		appendTreeNode(root.Left, index+1)
		appendTreeNode(root.Right, index+1)
	}
}
