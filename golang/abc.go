package main

import (
	"fmt"
	"sync"
)

// 编写一个程序，启动三个线程，三个线程的ID分别是A，B，C；，每个线程将自己的ID值在屏幕上打印5遍，打印顺序是ABCABC…
func main() {
	cha := make(chan int)
	chb := make(chan int)
	chc := make(chan int)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	wg.Add(1)
	wg.Add(1)

	go printStr(wg, cha, chb, "A")
	go printStr(wg, chb, chc, "B")
	go printStr(wg, chc, cha, "C")

	cha <- 1
	wg.Wait()
}

func printStr(wg *sync.WaitGroup, cha chan int, chb chan int, str string) {
	defer wg.Done()
	i := 0
	for {

		select {
		case <-cha:
			i++
			fmt.Println(str)
			if i >= 5 {
				close(chb)
				return
			}
			chb <- i
		}
	}
}
