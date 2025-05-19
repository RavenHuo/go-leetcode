package main

import "fmt"

// 将正确的数字放到对应的 i-1 下标上
func firstMissingPositive(nums []int) int {

	// 把每个数字 n 放到 nums[n-1] 的位置
	for i := 0; i < n; i++ {
		for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}
	for i, n := range nums {
		if i+1 != n {
			return i + 1
		}
	}
	return len(nums) + 1
}
func main() {
	fmt.Println(firstMissingPositive([]int{3, 4, -1, 1}))
}
