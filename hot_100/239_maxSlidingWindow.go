package main

import (
	"fmt"
)

// 单调队列
func maxSlidingWindow(nums []int, k int) []int {
	if k == 1 {
		return nums
	}
	result := make([]int, 0)
	// 有序的下标,最大的排前面
	indexArr := make([]int, 0)
	for i, n := range nums {
		if len(indexArr) > 0 {
			// 将下标数组中比自己小的元素清理掉
			for len(indexArr) > 0 && nums[indexArr[len(indexArr)-1]] < n {
				indexArr = indexArr[:len(indexArr)-1]
			}
		}
		indexArr = append(indexArr, i)
		// 再判断头部需不需要更新，如果第一个一直不在范围内，需要一直踢出
		for indexArr[0] < i-k+1 {
			indexArr = indexArr[1:]
		}
		if i-k+1 >= 0 {
			result = append(result, nums[indexArr[0]])
		}
	}
	return result
}

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, 1, 2, 0, 5}, 3))
}
