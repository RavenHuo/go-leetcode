/**
 * @Author raven
 * @Description
 * @Date 2022/3/12
 **/
package main

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
