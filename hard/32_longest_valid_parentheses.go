/**
 * @Author raven
 * @Description
 * @Date 2022/11/15
 **/
package main

import (
	"fmt"
)

func longestValidParentheses(s string) int {
	strStack := newStack(len(s))
	result := 0
	cur := 0
	// 初始化位置结果下标
	pos := make([]int, len(s))
	flag := false
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "(" {
			strStack.Push(string(s[i]))
			flag = false
		} else {
			popStr, _ := (strStack.Pop()).(string)
			if popStr == "(" {
				if flag {
					if i-1 >= 0 {
						// 去前n个结果获取是否 是连续的
						cur = pos[i-1] + 2
					} else {
						cur = 2
					}
				} else {
					if i-2 >= 0 {
						cur = pos[i-2] + 2
					} else {
						cur = 2
					}
				}
				// 连续相加
				if i-cur > 0 {
					cur += pos[i-cur]
				}
			} else {
				cur = 0
			}
			pos[i] = cur
			if cur > result {
				result = cur
			}
			flag = true
		}
	}
	return result
}

type Stack struct {
	cap  int
	data []interface{}
}

// Push 进栈
func (s *Stack) Push(item interface{}) {
	if s.Len()+1 > s.cap {
		panic("reach cap")
	}
	s.data = append(s.data, item)
}

func (s *Stack) Len() int {
	return len(s.data)
}

//Pop 出栈
func (s *Stack) Pop() interface{} {
	if len(s.data) > 0 {
		result := s.data[s.Len()-1]
		s.data = append(s.data[:s.Len()-1], s.data[s.Len()-1+1:]...)
		return result
	}
	return nil
}

func (s *Stack) Full() bool {
	return s.Len()+1 >= s.cap
}
func newStack(cap int) *Stack {
	return &Stack{cap: cap, data: make([]interface{}, 0, cap)}
}

//
func main() {
	fmt.Println(longestValidParentheses(")()())"))
}
