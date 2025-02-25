package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 21. 合并两个有序链表
// 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
// 首先这两个 链表是有序的
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head *ListNode
	var tail *ListNode
	for {
		if list1 == nil && list2 == nil {
			break
		}
		if list1 == nil {
			tail.Next = list2
			break
		}
		if list2 == nil {
			tail.Next = list1
			break
		}
		v1 := list1.Val
		v2 := list2.Val
		v := 0
		if v1 <= v2 {
			v = v1
			list1 = list1.Next
		} else {
			v = v2
			list2 = list2.Next
		}
		node := &ListNode{
			Val: v,
		}
		if head == nil {
			head = node
			tail = head
		} else {
			tail.Next = node
			tail = tail.Next
		}
	}
	return head
}

func main() {

}
