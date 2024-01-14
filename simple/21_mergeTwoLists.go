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
	r := &ListNode{}
	c := r
	for list1 != nil || list2 != nil {
		if list1 == nil {
			c.Next = list2
			break
		}
		if list2 == nil {
			c.Next = list1
			break
		}
		if list1.Val <= list2.Val {
			c.Next = &ListNode{Val: list1.Val, Next: nil}
			c = c.Next
			list1 = list1.Next
		} else {
			c.Next = &ListNode{Val: list2.Val, Next: nil}
			c = c.Next
			list2 = list2.Next
		}
	}
	return r.Next
}

func main() {

}
