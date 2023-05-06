/**
 * @Author raven
 * @Description
 * @Date 2022/11/6
 **/
package data_struct

import (
	"container/list"
	"sync"
)

type WxPlay struct {
	mutex    *sync.Mutex
	safeList *SafeList
}

type SafeList struct {
	list  *list.List
	mutex *sync.Mutex
}

func InitSafeList() *SafeList {
	return &SafeList{
		list:  list.New(),
		mutex: &sync.Mutex{},
	}
}

// 新增一个节点在队尾
func (s *SafeList) Push(t interface{}) {
	s.mutex.Lock()
	rwLock := &sync.Map{}
	defer s.mutex.Unlock()
	s.list.PushBack(t)
}

// 从队列头部获取一个接点
func (s *SafeList) Pop() interface{} {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	element := s.list.Front()
	if element != nil {
		s.list.Remove(element)
		return element.Value
	}
	return nil
}
func NewWxPlay() *WxPlay {
	return &WxPlay{
		mutex:    &sync.Mutex{},
		safeList: InitSafeList(),
	}
}

func (play *WxPlay) play() {
	play.mutex.Lock()
	defer play.mutex.Unlock()
	element := play.safeList.Pop()
	if element != nil {
		// 播放
	}
}
func (play *WxPlay) listen(t interface{}) {
	play.safeList.Push(t)
}
