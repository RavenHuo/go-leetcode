package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, tail *ListNode
	var flag bool
	for l1 != nil || l2 != nil || flag {
		v1 := 0
		v2 := 0
		v := 0
		extra := 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		if flag {
			extra = 1
		}
		if v1+v2+extra >= 10 {
			flag = true
			v = v1 + v2 + extra - 10
		} else {
			flag = false
			v = v1 + v2 + extra
		}
		if head == nil {
			head = &ListNode{Val: v}
			tail = head
		} else {
			tail.Next = &ListNode{Val: v}
			tail = tail.Next
		}
	}
	return head
}
