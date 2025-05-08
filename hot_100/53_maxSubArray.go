package main

import (
	"fmt"
	"sync"
	"time"
)

// 动态规划
// 53. 最大子数组和
// 状态转移方程：当前元素重新开始或者跟之前的合并为一个子数组
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	resultList := make([]int, len(nums))
	resultList[0] = nums[0]
	m := nums[0]
	for i := 1; i <= len(nums)-1; i++ {
		n := nums[i]
		resultList[i] = maxNum(n, resultList[i-1]+n)
		if resultList[i] > m {
			m = resultList[i]
		}
	}

	return m
}

func maxNum(n, m int) int {
	if n > m {
		return n
	}
	return m
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(time.Second)
	}()
	go func() {
		wg.Wait()
		fmt.Println("go wg wait ")
	}()
	fmt.Println("main wg wait ")
	time.Sleep(time.Second)
}
