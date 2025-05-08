package main

import (
	"fmt"
	"time"
)

// 交替打印1，2，3
func main() {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	go handleCh1(ch1, ch2)
	go handleCh2(ch1, ch2)
	ch1 <- struct{}{}
	time.Sleep(time.Second * 5)
}

var n int

func handleCh1(ch1, ch2 chan struct{}) {
	for {
		select {
		case <-ch1:
			if n >= 10 {
				close(ch2)
				return
			} else {
				fmt.Println(n)
				n = n + 1
				ch2 <- struct{}{}
			}
		}
	}
}
func handleCh2(ch1, ch2 chan struct{}) {
	for {
		select {
		case <-ch2:
			fmt.Println(n)
			if n >= 10 {
				close(ch1)
				return
			} else {
				n = n + 1
				ch1 <- struct{}{}
			}
		}
	}
}
