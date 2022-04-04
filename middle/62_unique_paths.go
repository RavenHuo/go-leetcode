/**
 * @Author raven
 * @Description
 * @Date 2022/4/4
 **/
package main

func swapPairs(head *ListNode) *ListNode {
	i := 1
	headNode := head
	lastNode := head
	for {
		if headNode == nil {
			break
		}
		if i%2 == 1 {
			lastNode = headNode
		} else {
			// 交换值
			lastNodeVal := lastNode.Val
			lastNode.Val = headNode.Val
			headNode.Val = lastNodeVal
		}
		headNode = headNode.Next
		i++
	}

	return head
}
