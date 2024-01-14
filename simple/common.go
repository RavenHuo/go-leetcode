/**
 * @Author raven
 * @Description
 * @Date 2022/3/12
 **/
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

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
	l1 := buildListNodeByArr([]int{1, 2, 4})
	l2 := buildListNodeByArr([]int{1, 3, 4})
	r := mergeTwoLists(l1, l2)
	fmt.Println(r)
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
