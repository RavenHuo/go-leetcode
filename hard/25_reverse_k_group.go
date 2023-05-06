/**
 * @Author raven
 * @Description
 * @Date 2022/11/14
 **/
package main

func reverseKGroup(head *ListNode, k int) *ListNode {
	listLen := 0
	item := head
	for item != nil {
		listLen++
		item = item.Next
	}
	if listLen == 0 {
		return head
	}
	// 翻转多少次
	cycle := listLen / k
	if cycle == 0 {
		return head
	}
	round := 0
	i := 0

	var result *ListNode
	item = result
	var stack *Stack
	for head != nil {
		if round == cycle {
			break
		}
		if i == 0 {
			stack = newStack(k)
		}
		if i < k {
			stack.Push(head.Val)
			i++
		}
		if i == k {
			i = 0
			round++
			for stack.Len() != 0 {
				s, _ := (stack.Pop()).(int)
				if result == nil {
					result = &ListNode{}
					item = result
				}
				item.Val = s

				if round == cycle && stack.Len() == 0 {
					item.Next = head.Next
					break
				} else {
					item.Next = &ListNode{}
					item = item.Next
				}
			}

		}
		head = head.Next
	}
	return result
}
