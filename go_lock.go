/**
 * @Author raven
 * @Description
 * @Date 2022/5/15
 **/
package main

import (
	"sync"
)

func main() {
	// 校验 写锁 加锁是否需要等到读锁释放
	//rwLock := sync.RWMutex{}
	//rwLock.RLock()
	//rwLock.Lock()
	//rwLock.Unlock()
	//rwLock.RUnlock()
	wp := sync.WaitGroup{}
	once := sync.Once{}
	m := sync.Map{}
	m.Delete()
}
