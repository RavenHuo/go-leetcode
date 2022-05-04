/**
 * @Author raven
 * @Description
 * @Date 2022/4/25
 **/
package main

import (
	"fmt"
)

// 观察者 interface
// 注册观察者数组，当被观察对象修改时，通知观察者
type Observer interface {
	Notify(date string) error
}

// 观察者
type ChangeObServer struct{}

func (f *ChangeObServer) Notify(date string) error {
	fmt.Println("chaneg name :" + date)
	return nil
}

type Subject struct {
	Observers []Observer
	Name      string
}

func InitSubject(n string) (s *Subject) {
	return &Subject{
		Name: n,
		Observers: []Observer{
			&ChangeObServer{}},
	}
}

func (s *Subject) Update(n string) {
	s.Name = n
	for _, o := range s.Observers {
		o.Notify(n)
	}
}

func main() {
	subject := InitSubject("raven")
	subject.Update("riven")
}
