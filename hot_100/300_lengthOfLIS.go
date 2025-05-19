package main

// 当前的 最长递增子序列 = 1+ 之前最长的递增至序列 （如果当前的数是最大的话） dp
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	for i := range len(nums) {
		dp[i] = 1
	}
	maxLen := 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if dp[j]+1 > dp[i] {
					// 之前最长的加上自己等于当前最长的
					dp[i] = dp[j] + 1
				}
			}
		}
		if dp[i] > maxLen {
			maxLen = dp[i]
		}
	}
	return maxLen
}
