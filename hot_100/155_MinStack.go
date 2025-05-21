package main

import "fmt"

type MinStack struct {
	data []int
	min  int
}

func Constructor() MinStack {
	return MinStack{
		data: make([]int, 0),
		min:  0,
	}
}

func (this *MinStack) Push(val int) {
	if len(this.data) == 0 {
		this.min = val
	} else {
		if this.min > val {
			this.min = val
		}
	}
	this.data = append(this.data, val)

}

func (this *MinStack) Pop() {
	top := this.data[len(this.data)-1]
	this.data = this.data[:len(this.data)-1]
	// 重置
	if top == this.min {
		if len(this.data) == 0 {
			this.min = 0
		} else {
			m := this.data[0]
			for i := 1; i < len(this.data); i++ {
				if this.data[i] < m {
					m = this.data[i]
				}
			}
			this.min = m
		}
	}
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}
func main() {
	s := Constructor()
	s.Push(-2)
	s.Push(0)
	s.Push(-3)
	fmt.Println(s.GetMin())
	s.Pop()
	s.Top()
	fmt.Println(s.GetMin())
}
