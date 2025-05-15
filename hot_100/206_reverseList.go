package main

func reverseList(head *ListNode) *ListNode {
	// 头
	var result *ListNode
	for head != nil {
		cur := head
		head = cur.Next

		cur.Next = result
		result = cur
	}
	return result
}
func main() {
	reverseList(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}})
}
