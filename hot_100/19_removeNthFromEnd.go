package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	tail := head
	count := 0
	m := make(map[int]*ListNode)
	for tail != nil {
		count++
		m[count] = tail
		tail = tail.Next
	}
	if n > count {
		return nil
	}
	removeBeforeIndex := count - n
	removeAfterIndex := count - n + 2
	if removeAfterIndex < 0 {
		return nil
	}
	// 移除了头节点
	if removeBeforeIndex > 0 {
		m[removeBeforeIndex].Next = m[removeAfterIndex]
	} else {
		return m[removeAfterIndex]
	}
	return head
}
