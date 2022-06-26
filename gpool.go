package main

import (
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
	//for i := 0; i < delta; i++ {
	//	p.queue <- 1
	//}
	//for i := 0; i > delta; i-- {
	//	<-p.queue
	//}
	p.wg.Add(delta)
}

func (p *pool) Done() {
	//<-p.queue
	p.wg.Done()
}

func (p *pool) Wait() {
	p.wg.Wait()
}
func main() {
	p := NewPool(1)
	p.Add(1)
	go func() {
		time.Sleep(10 * time.Minute)
		p.Done()
	}()
	p.Add(1)
	p.Wait()
}
