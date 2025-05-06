package main

// 就是判断i能不能追上j， i每一次走一步，j每两次走一步
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	iNode := head.Next
	jNode := head
	c := 0
	for (iNode != nil) && (jNode != nil) {
		if iNode == jNode {
			return true
		}
		if c%2 == 1 {
			jNode = jNode.Next
		}
		iNode = iNode.Next
		c++
	}
	return false
}
