/**
 * @Author raven
 * @Description
 * @Date 2022/3/12
 **/
package main

import (
	"log"
)

// 2. 两数相加
// 给你两个非空 的链表，表示两个非负的整数。它们每位数字都是按照逆序的方式存储的，并且每个节点只能存储一位数字。
//
//请你将两个数相加，并以相同形式返回一个表示和的链表。
//
//你可以假设除了数字 0 之外，这两个数都不会以 0开头。
//
//来源：力扣（LeetCode） 2. 两数相加
//链接：https://leetcode-cn.com/problems/add-two-numbers
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}
	var result *ListNode
	var header *ListNode
	// 值
	val := 0
	// 判断是否进一
	flag := false
	for {
		if l1 == nil && l2 == nil && !flag {
			break
		}
		if flag == true {
			val += 1
			flag = false
		}
		if l1 != nil {
			val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val += l2.Val
			l2 = l2.Next
		}
		if val >= 10 {
			val -= 10
			flag = true
		}
		node := &ListNode{Val: val}
		node.Next = nil
		if header == nil {
			header = node
			result = header
		} else {
			header.Next = node
			header = node
		}
		val = 0
	}

	return result
}

func addTwoNumbersTest() {
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	result := addTwoNumbers(l1, l2)
	log.Printf("addTwoNumbers :+%v", result)
}
