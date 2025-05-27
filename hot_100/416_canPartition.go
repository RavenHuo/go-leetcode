package main

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	dp := make([]bool, target+1)
	dp[0] = true
	for _, num := range nums {
		for j := target; j >= num; j-- {
			// 就是能不能达到 dp[j] 能达到的话就是true
			dp[j] = dp[j] || dp[j-num]
		}
	}
	return dp[target]
}
