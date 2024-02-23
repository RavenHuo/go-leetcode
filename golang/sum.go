package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 编写10个线程，第一个线程从1加到10，第二个线程从11加20…第十个线程从91加到100，最后再把10个线程结果相加。
func main() {
	wg := &sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(n int) {
			defer wg.Done()
			nums := make([]int, 0)
			for j := 0; j < 10; j++ {
				nums = append(nums, n*10+j)
			}
			sum(nums)
		}(i)
	}
	wg.Wait()
	fmt.Println(s.Load())
}

var s = atomic.Int32{}

func sum(nums []int) {
	for _, v := range nums {
		s.Add(int32(v))
	}
}
