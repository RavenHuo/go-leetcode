package main

import "container/list"

type LRUCache struct {
	m        map[int]*list.Element
	l        *list.List
	capacity int
}

type LruCacheEntry struct {
	k int
	v int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		m:        make(map[int]*list.Element),
		l:        &list.List{},
	}
}

func (this *LRUCache) Get(key int) int {
	e, ok := this.m[key]
	if ok {
		this.l.MoveToFront(e)
		return e.Value.(*LruCacheEntry).v
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	e, ok := this.m[key]
	if ok {
		entry := e.Value.(*LruCacheEntry)
		entry.v = value
		this.l.MoveToFront(e)
		return
	}
	entry := &LruCacheEntry{
		k: key,
		v: value,
	}
	e = &list.Element{
		Value: entry,
	}
	if len(this.m) == this.capacity {
		b := this.l.Back()
		if b != nil {
			backEntry := b.Value.(*LruCacheEntry)
			delete(this.m, backEntry.k)
			this.l.Remove(b)
		}

	}
	this.m[key] = e
	this.l.MoveToFront(e)
	return
}

func main() {
	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(1, 2)
	lru.Put(2, 2)
	lru.Put(2, 1)
	lru.Put(3, 1)
}
