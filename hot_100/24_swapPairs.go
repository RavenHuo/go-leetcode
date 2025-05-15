package main

func swapPairs(head *ListNode) *ListNode {
	c := 0
	tail := head
	var pre *ListNode
	for tail != nil {
		// 交换
		if c%2 == 0 && tail.Next != nil {
			if pre == nil {
				cur := tail
				next := tail.Next
				nextNext := next.Next
				tail = next
				tail.Next = cur
				tail.Next.Next = nextNext
			} else {
				cur := tail
				next := tail.Next
				nextNext := next.Next
				tail = next
				tail.Next = cur
				tail.Next.Next = nextNext
				pre.Next = tail
			}
		}
		c++
		if pre == nil {
			head = tail
		}
		pre = tail
		tail = tail.Next
	}
	return head
}

func main() {
	swapPairs(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}})
}
