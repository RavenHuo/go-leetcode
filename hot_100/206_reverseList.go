package main

func reverseList(head *ListNode) *ListNode {
	// 头
	var result *ListNode
	for head != nil {
		v := head.Val
		node := &ListNode{Val: v, Next: result}
		result = node
		head = head.Next
	}

	return result
}
