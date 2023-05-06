/**
 * @Author raven
 * @Description
 * @Date 2022/11/14
 **/
package data_struct

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
