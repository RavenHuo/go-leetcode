package main

import "container/list"

type LRUCache struct {
	capacity int
	m        map[int]*list.Element
	l        *list.List
}

type LruEntry struct {
	k int
	v int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		m:        make(map[int]*list.Element),
		l:        list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
	e, ok := this.m[key]
	if ok {
		this.l.MoveToFront(e)
		return e.Value.(*LruEntry).v
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	e, ok := this.m[key]
	if ok {
		entry := e.Value.(*LruEntry)
		entry.k = key
		entry.v = value
		e.Value = entry
		this.l.MoveToFront(e)
		return
	}
	// 不存在
	// 超过容量就要驱逐
	if len(this.m)+1 > this.capacity {
		back := this.l.Back()
		if back != nil {
			backEntry := back.Value.(*LruEntry)
			delete(this.m, backEntry.k)
			this.l.Remove(back)
		}
	}
	entry := &LruEntry{
		k: key,
		v: value,
	}
	e = this.l.PushFront(entry)
	this.m[key] = e
}
