package main

import "container/list"

type Lru struct {
	size int
	m    map[int]*list.Element
	l    *list.List
}
type LruEntry struct {
	value int
	key   int
}

func NewLru(s int) *Lru {
	return &Lru{
		size: s,
		m:    make(map[int]*list.Element),
		l:    list.New(),
	}
}

func (l *Lru) Get(key int) int {
	element, ok := l.m[key]
	if !ok {
		return -1
	}
	entry := element.Value.(*LruEntry)
	l.l.MoveToFront(element)
	return entry.value
}

func (l *Lru) Put(key, value int) {
	element, ok := l.m[key]
	if ok {
		entry := element.Value.(*LruEntry)
		entry.value = value
		element.Value = entry
		l.l.MoveToFront(element)
		l.m[key] = element
		return
	} else {
		entry := &LruEntry{value: value, key: key}
		if len(l.m) > l.size {
			back := l.l.Back()
			backEntry := back.Value.(*LruEntry)
			delete(l.m, backEntry.key)
			l.l.Remove(back)
		}
		element = l.l.PushFront(entry)
		l.m[key] = element
	}
}
