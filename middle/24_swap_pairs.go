/**
 * @Author raven
 * @Description
 * @Date 2022/4/4
 **/
package main

//func main() {
//	l := rotateRight(buildListNodeByArr([]int{1,2,3,4,5}), 2)
//	for {
//		if l == nil {
//			break
//		}
//		fmt.Println(l.Val)
//		l = l.Next
//	}
//}

// 61. 旋转链表 给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。
func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil {
		return head
	}
	headNode := head
	// 链表长度
	listLen := 0
	// 构造循环链表
	for {
		if headNode.Next == nil {
			listLen++
			headNode.Next = head
			break
		}
		headNode = headNode.Next
		listLen++
	}

	var result *ListNode
	var resultNode *ListNode
	// 重新开始
	headNode = head
	n := listLen - (k % listLen)
	for {
		if n == 0 && listLen == 0 {
			break
		}
		if n > 0 {
			headNode = headNode.Next
			n--
		}
		if n == 0 && listLen > 0 {
			node := &ListNode{Val: headNode.Val}
			if resultNode == nil {
				resultNode = node
				result = resultNode
			} else {
				resultNode.Next = node
				resultNode = node
			}
			listLen--
			headNode = headNode.Next
		}
	}

	return result
}
