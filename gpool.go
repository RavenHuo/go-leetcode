package main

import (
	"fmt"
	"sync"
	"time"
)

type Pool interface {
	Add(delta int)
	Done()
	Wait()
}

type pool struct {
	queue chan int
	wg    *sync.WaitGroup
}

func NewPool(size int) *pool {
	if size <= 0 {
		size = 1
	}
	return &pool{
		queue: make(chan int, size),
		wg:    &sync.WaitGroup{},
	}
}

func (p *pool) Add(delta int) {
	for i := 0; i < delta; i++ {
		p.queue <- 1
	}
	p.wg.Add(delta)
}

func (p *pool) Done() {
	<-p.queue
	p.wg.Done()
}

func (p *pool) Wait() {
	p.wg.Wait()
}
func main() {
	p := NewPool(1)
	p.Add(1)
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("goroutine 1")
		p.Done()
	}()
	p.Add(1)
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("goroutine 2")
		p.Done()
	}()
	p.Wait()
}
