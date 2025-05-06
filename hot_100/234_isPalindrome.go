package main

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}
	reverse := reverseList(head)
	for (head != nil) || (reverse != nil) {
		if head == nil {
			return false
		}
		if reverse == nil {
			return false
		}
		if head.Val != reverse.Val {
			return false
		}
		if head != nil {
			head = head.Next
		}
		if reverse != nil {
			reverse = reverse.Next
		}
	}
	return true
}
