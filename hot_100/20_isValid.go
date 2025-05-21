package main

func isValid(s string) bool {
	charList := make([]byte, 0)
	for _, c := range []byte(s) {
		strC := string(c)
		if strC == "(" || strC == "{" || strC == "[" {
			charList = append(charList, c)
		} else {
			if len(charList) == 0 {
				return false
			}
			if strC == "}" && string(charList[len(charList)-1]) != "{" {
				return false
			}
			if strC == ")" && string(charList[len(charList)-1]) != "(" {
				return false
			}
			if strC == "]" && string(charList[len(charList)-1]) != "[" {
				return false
			}
			charList = charList[:len(charList)-1]
		}
	}
	if len(charList) != 0 {
		return false
	}
	return true
}

type Stack struct {
	data []int
}

func NewStack() *Stack {
	return &Stack{data: make([]int, 0)}
}

func (s *Stack) Push(n int) {
	s.data = append(s.data, n)
}

func (s *Stack) Pop() int {
	if len(s.data) == 0 {
		return -1
	}
	n := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return n
}
