package main

import (
	"fmt"
	"sync"
)

// WorkerPool 控制并发执行的协程数量
type WorkerPool struct {
	size  int
	jobCh chan func()
	wg    *sync.WaitGroup
}

func InitWorkerPool(size int) *WorkerPool {
	return &WorkerPool{
		size:  size,
		jobCh: make(chan func(), size),
		wg:    &sync.WaitGroup{},
	}
}

func (w *WorkerPool) Submit(job func()) {
	w.jobCh <- job
	w.wg.Add(1)
}

// 真正的执行者
func (w *WorkerPool) worker() {
	for {
		select {
		case job := <-w.jobCh:
			func(j func()) {
				defer func() {
					w.wg.Done()
					r := recover()
					if r != nil {
						fmt.Printf("recover r:%s \n", r)
					}
				}()
				j()
			}(job)
		}
	}
}

// Start 开协程来处理
func (w *WorkerPool) Start() {
	for i := 0; i < w.size; i++ {
		go w.worker()
	}
}

func (w *WorkerPool) Wait() {
	w.wg.Wait()
	close(w.jobCh)
	return
}
