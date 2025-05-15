package main

// 每k个一组反转链表
// 断链再接回去
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	c := 0
	tail := head
	var pre *ListNode
	var res *ListNode
	var cur *ListNode
	var curHead *ListNode
	for tail != nil {
		c++
		node := &ListNode{Val: tail.Val}
		if cur == nil {
			cur = node
			curHead = cur
		} else {
			cur.Next = node
			cur = cur.Next
		}
		if c%k == 0 {
			if pre == nil {
				pre = reverseList2(curHead)
				res = pre
				for pre.Next != nil {
					pre = pre.Next
				}
			} else {
				pre.Next = reverseList2(curHead)
				for pre.Next != nil {
					pre = pre.Next
				}
			}
			// 清空cur
			cur = nil
			curHead = nil
		}
		tail = tail.Next
	}
	pre.Next = curHead
	return res
}

func reverseList2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var node *ListNode
	for head != nil {
		next := &ListNode{Val: head.Val, Next: node}
		node = next
		head = head.Next
	}
	return node
}

func main() {
	reverseKGroup(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}, 2)
}
