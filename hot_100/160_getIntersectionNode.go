package main

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	rmap := make(map[*ListNode]struct{})
	for headA != nil {
		rmap[headA] = struct{}{}
		headA = headA.Next
	}
	for headB != nil {
		_, ok := rmap[headB]
		if ok {
			return headB
		}
		headB = headB.Next
	}
	return nil
}
