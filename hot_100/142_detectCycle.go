package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 快慢指针：
/*
   1. 快指针每次走两步，慢指针每次走一步
   2. 走到第一次相遇后，慢指针再从起点开始走，快指针也变成慢指针走，再次相遇的节点就是环的入口
*/

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	iNode := head
	jNode := head

	for iNode != nil && iNode.Next != nil {
		iNode = iNode.Next.Next
		jNode = jNode.Next
		if iNode == jNode {
			jNode = iNode
			iNode = head
			for iNode != jNode {
				iNode = iNode.Next
				jNode = jNode.Next
			}
			return iNode
		}
	}
	return nil
}

func main() {
	fmt.Println("hello")

	n := &ListNode{Val: 1}
	n.Next = &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: n}}}
	l := n
	fmt.Println(l.Val)
	fmt.Println(detectCycle(l))
}
