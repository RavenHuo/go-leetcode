package main

import (
	"container/list"
	"sync"
)

type LRUCache struct {
	lock    *sync.RWMutex
	cap     int
	itemMap map[string]*list.Element
	list    *list.List
}

func NewLruCache(cap int) *LRUCache {
	return &LRUCache{
		lock:    &sync.RWMutex{},
		cap:     cap,
		itemMap: make(map[string]*list.Element),
		list:    list.New(),
	}
}

func (c *LRUCache) add(key string, value interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()
	item, ok := c.itemMap[key]
	if ok {
		item.Value = value
		c.list.MoveToFront(item)
		c.itemMap[key] = item
	} else {
		item = &list.Element{
			Value: value,
		}
		if c.list.Len() > c.cap {
			c.list.Remove(c.list.Back())
		}
		c.itemMap[key] = item
		c.list.PushFront(item)
	}
}
func (c *LRUCache) Get(key string) interface{} {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.itemMap[key]
}

func (c *LRUCache) Remove(key string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	item, ok := c.itemMap[key]
	if !ok {
		return
	}
	c.list.Remove(item)
	delete(c.itemMap, key)
}
