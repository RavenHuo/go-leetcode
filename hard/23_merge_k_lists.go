/**
 * @Author raven
 * @Description
 * @Date 2022/4/5
 **/
package main

// 分治法 先合并两个listNode，然后再逐个合并
// 合并多个list
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return merge(lists, 0, len(lists)-1)
}

// for 循环遍历 合并
func forMerge(lists []*ListNode) *ListNode {
	var result *ListNode
	for _, list := range lists {
		if result == nil {
			result = mergeTwoList(list, nil)
		} else {
			result = mergeTwoList(result, list)
		}
	}
	return result
}

// 二分法 两两合并，left，right，mid
func merge(lists []*ListNode, left, right int) *ListNode {
	if left == right {
		return lists[right]
	}
	mid := (left + right) / 2
	leftNode := merge(lists, left, mid)
	rightNode := merge(lists, mid+1, right)
	return mergeTwoList(leftNode, rightNode)
}

// 合并两个list
func mergeTwoList(left *ListNode, right *ListNode) *ListNode {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	var result *ListNode
	var resultNode *ListNode
	for {
		if left == nil && right == nil {
			break
		}
		var val int
		if left != nil && right != nil {
			if left.Val > right.Val {
				val = right.Val
				right = right.Next
			} else {
				val = left.Val
				left = left.Next
			}
		} else if left != nil && right == nil {
			val = left.Val
			left = left.Next
		} else if right != nil && left == nil {
			val = right.Val
			right = right.Next
		}
		node := &ListNode{Val: val}
		if resultNode == nil {
			resultNode = node
			result = resultNode
		} else {
			resultNode.Next = node
			resultNode = node
		}
	}
	return result
}

//func main() {
//	lists := make([]*ListNode, 0)
//	lists = append(lists, buildListNodeByArr([]int{1, 4, 5}))
//	lists = append(lists, buildListNodeByArr([]int{1, 3, 4}))
//	lists = append(lists, buildListNodeByArr([]int{2, 6}))
//	l := mergeKLists(lists)
//	for {
//		if l == nil {
//			break
//		}
//		fmt.Println(l.Val)
//		l = l.Next
//	}
//}
