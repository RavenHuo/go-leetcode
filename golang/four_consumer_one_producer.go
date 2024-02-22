package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go produce(ch, wg)
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go consumer(ch, wg)
	}
	wg.Wait()
	fmt.Println("end")
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for item := range ch {
		fmt.Println(item)
	}
}

func produce(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(ch)
	for i := 0; i < 100; i++ {
		ch <- i
	}
}
