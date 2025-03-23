package main

import (
	"fmt"
	"sync"
)

// 两个协程交错打印1,2,3,4,5,6,7,8,9,10
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go handleCh1(wg, ch1, ch2)
	go handlerCh2(wg, ch2, ch1)
	ch1 <- 1
	wg.Wait()
}

func handlerCh2(wg *sync.WaitGroup, ch2 chan int, ch1 chan int) {
	wg.Done()
	for {
		select {
		case n := <-ch2:
			if n > 10 {
				close(ch2)
				return
			}
			fmt.Println(n)
			ch1 <- n + 1
		}
	}
}

func handleCh1(wg *sync.WaitGroup, ch1 chan int, ch2 chan int) {
	defer wg.Done()
	for {
		select {
		case n := <-ch1:
			if n > 10 {
				close(ch1)
				ch2 <- n + 1
				return
			}
			fmt.Println(n)
			ch2 <- n + 1
		}
	}
}
