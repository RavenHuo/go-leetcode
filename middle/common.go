/**
 * @Author raven
 * @Description
 * @Date 2022/3/12
 **/
package main

import (
	"fmt"
)

func buildListNodeByArr(list []int) *ListNode {
	var header *ListNode
	var result *ListNode
	for _, val := range list {
		node := &ListNode{Val: val}
		node.Next = nil
		if header == nil {
			header = node
			result = header
		} else {
			header.Next = node
			header = node
		}
	}
	return result
}

func main() {
	h := buildListNodeByArr([]int{1, 2, 3, 4})
	reverseLinkedList(h)
	fmt.Printf("%+v", h)
}

func reverseLinkedList(head *ListNode) {
	var nextPoint *ListNode
	point := head

	for point != nil {
		v := point.Val
		node := &ListNode{Val: v, Next: nextPoint}
		nextPoint = node
		point = point.Next
	}
	head = nextPoint
}

func app(l1 *ListNode) *ListNode {
	var header *ListNode
	var tail *ListNode
	for {
		if l1 == nil {
			break
		}
		val := &ListNode{
			Val: l1.Val,
		}
		if header == nil {
			header = val
			tail = val
		} else {
			tail.Next = val
			tail = tail.Next
		}
	}
	return header
}
