/**
 * @Author raven
 * @Description
 * @Date 2022/6/14
 **/
package main

func reorderList(head *ListNode) {
	valueList := make([]int, 0)

	item := head
	for item != nil {
		valueList = append(valueList, item.Val)
		item = item.Next
	}
	j := len(valueList) - 1
	i := 0
	item = head
	n := 0
	for i <= j {
		if n%2 == 1 {
			item.Val = valueList[j]
			j--
		} else {
			item.Val = valueList[i]
			i++
		}
		item = item.Next
		n++
	}
}
