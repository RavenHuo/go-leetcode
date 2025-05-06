package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var tail *ListNode
	var head *ListNode
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	firstVal := 0
	if list1.Val < list2.Val {
		firstVal = list1.Val
	} else {
		firstVal = list2.Val
	}
	head = &ListNode{Val: firstVal}
	tail = head

	for list1 != nil || list2 != nil {
		var v *ListNode
		if list1 != nil && list2 != nil {
			v1 := list1.Val
			v2 := list2.Val
			if v1 <= v2 {
				v = list1
				list1 = list1.Next
			} else {
				v = list2
				list2 = list2.Next
			}
		} else if list1 != nil && list2 == nil {
			v = list1
			list1 = list1.Next
		} else if list1 == nil && list2 != nil {
			v = list2
			list2 = list2.Next
		}
		fmt.Println(v)
		v.Next = nil
		tail.Next = v
	}
	return head
}

func main() {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}
	l2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	mergeTwoLists(l1, l2)
}
