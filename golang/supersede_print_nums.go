package main

import (
	"fmt"
	"sync"
)

// 两个协程相互打印1-100
// ch1 print 单数
// ch2 print双数
func main() {
	wg := &sync.WaitGroup{}
	c := 0
	// 通过ch通知
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ch1:
				if c >= 100 {
					// 通知ch2关闭
					close(ch2)
					return
				}
				c++
				fmt.Println(c)
				ch2 <- struct{}{}
			}
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {

			select {
			case <-ch2:
				if c >= 100 {
					// 通知ch1关闭
					close(ch1)
					return
				}
				c++
				fmt.Println(c)
				ch1 <- struct{}{}
			}
		}
	}()
	ch1 <- struct{}{}
	wg.Wait()

}
