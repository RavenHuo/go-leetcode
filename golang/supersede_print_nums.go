package main

import (
	"fmt"
	"sync"
)

// 两个协程相互打印1-100
// ch1 print 单数
// ch2 print双数
func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)
	sum := 0
	wg := &sync.WaitGroup{}
	wg.Add(1)
	wg.Add(1)
	go func(*sync.WaitGroup) {
		defer wg.Done()
		for {
			_ = <-ch1
			fmt.Println(sum)
			sum++
			if sum > 100 {
				return
			}
			ch2 <- sum
		}
	}(wg)
	go func(*sync.WaitGroup) {
		defer wg.Done()
		for {
			_ = <-ch2
			fmt.Println(sum)
			sum++
			if sum > 100 {
				return
			}
			ch1 <- sum

		}
	}(wg)
	ch1 <- 1

	wg.Wait()
	fmt.Println("end")
}
