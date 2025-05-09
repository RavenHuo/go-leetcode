package main

import (
	"container/list"
	"sync"
)

type LRUCache struct {
	m     map[string]*list.Element
	l     *list.List
	count int // 容量
	lock  *sync.Mutex
}

func NewLru(count int) *LRUCache {
	return &LRUCache{
		m:     make(map[string]*list.Element),
		l:     list.New(),
		count: count,
		lock:  &sync.Mutex{},
	}
}
