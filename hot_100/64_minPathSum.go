package main

func minPathSum(grid [][]int) int {

	dp := make([][]int, len(grid))
	lSize := len(grid[0])
	for i := range len(grid) {
		dp[i] = make([]int, lSize)
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < lSize; j++ {
			if i == 0 && j == 0 {
				dp[0][0] = grid[0][0]
			} else if i-1 < 0 {
				dp[i][j] = dp[0][j-1] + grid[i][j]
			} else if j-1 < 0 {
				dp[i][j] = dp[i-1][0] + grid[i][j]
			} else {
				if dp[i-1][j] < dp[i][j-1] {
					dp[i][j] = dp[i-1][j] + grid[i][j]
				} else {
					dp[i][j] = dp[i][j-1] + grid[i][j]
				}
			}
		}
	}
	return dp[len(grid)-1][lSize-1]
}
