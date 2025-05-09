package main

// 动态规划
// 53. 最大子数组和
// 当前下标的最大和就是当前n+前
func maxSubArray(nums []int) int {
	arr := make([]int, len(nums))
	arr[0] = nums[0]
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		n := nums[i]
		arr[i] = maxNum(n, arr[i-1]+n)
		max = maxNum(arr[i], max)
	}
	return max
}

func maxNum(a, b int) int {
	if a > b {
		return a
	}
	return b
}
