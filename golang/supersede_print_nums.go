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
	count := 0
	ch1 := make(chan int)
	ch2 := make(chan int)
	wg.Add(1)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ch1:
				if count >= 100 {
					close(ch2)
					return
				}
				count++
				fmt.Println(count)
				ch2 <- 1
			}
		}
	}()

	go func() {
		defer wg.Done()
		for {
			select {
			case <-ch2:
				if count >= 100 {
					close(ch1)
					return
				}
				count++
				fmt.Println(count)
				ch1 <- 1
			}
		}
	}()
	ch1 <- 1
	wg.Wait()
}
