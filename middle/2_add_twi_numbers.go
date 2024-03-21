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
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
//
// 你可以假设除了数字 0 之外，这两个数都不会以 0开头。
//
// 来源：力扣（LeetCode） 2. 两数相加
// 链接：https://leetcode-cn.com/problems/add-two-numbers
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	var head *ListNode
	var tail *ListNode
	flag := false
	for {
		if l1 == nil && l2 == nil && !flag {
			break
		}
		v1 := 0
		v2 := 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		v := v1 + v2
		if flag == true {
			flag = false
			v += 1
		}
		if v >= 10 {
			v -= 10
			flag = true
		}
		node := &ListNode{Val: v}
		if head == nil {
			head = node
			tail = head
		} else {
			tail.Next = node
			tail = tail.Next
		}
	}
	return head
}

func addTwoNumbersTest() {
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	result := addTwoNumbers(l1, l2)
	log.Printf("addTwoNumbers :+%v", result)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	addTwoNumbersTest()
}
