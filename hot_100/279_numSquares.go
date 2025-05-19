package main

import "fmt"

func numSquares(n int) int {
	dp := make([]int, n+1)
	// 当前的平方数
	dpx := 1
	for i := 1; i <= n; i++ {
		dp[i] = i
	}
	for i := 1; i <= n; i++ {
		if i == (dpx+1)*(dpx+1) {
			dpx++
			dp[i] = 1
			continue
		}

		for j := dpx; j > 0; j-- {
			if i-j*j > 0 && dp[i] > dp[i-j*j]+1 {
				dp[i] = dp[i-j*j] + 1
			}
		}
	}
	return dp[n]
}

func main() {
	fmt.Println(numSquares(12))
}
