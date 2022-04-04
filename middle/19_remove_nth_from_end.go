/**
 * @Author raven
 * @Description
 * @Date 2022/4/3
 **/
package main

//func main() {
//	l := removeNthFromEnd(buildListNodeByArr([]int{2, 4, 3, 5, 6, 7}), 1)
//	for {
//		if l == nil {
//			break
//		}
//		fmt.Println(l.Val)
//		l = l.Next
//	}
//}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	listVal := make([]int, 0)
	for {
		if head == nil {
			break
		}
		listVal = append(listVal, head.Val)
		head = head.Next
	}
	removeIndex := len(listVal) - n
	var header *ListNode
	if removeIndex < 0 {
		return head
	}
	var result *ListNode
	for i, v := range listVal {
		if i == removeIndex {
			continue
		}
		node := &ListNode{Val: v}
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
