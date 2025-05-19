package main

import "fmt"

// 前缀和 就是全部加起来的和
func subarraySum(nums []int, k int) int {
	sumMap := make(map[int]int)
	sumMap[0] = 1
	sum := 0
	count := 0
	for _, n := range nums {
		sum += n
		countNum, ok := sumMap[sum-k]
		if ok {
			count += countNum
		}
		sumMap[sum] += 1
	}
	return count
}

func main() {
	fmt.Println(subarraySum([]int{1, 1, 1}, 2))
}
