package main

// dp[n][m] = dp[m-1][n] + dp[m]dp[n-1]
func uniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	paths := make([][]int, m)
	for k := range m {
		ints := make([]int, n)
		if k == 0 {
			for j := range n {
				ints[j] = 1
			}
		} else {
			ints[0] = 1
		}
		paths[k] = ints
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			paths[i][j] = paths[i-1][j] + paths[i][j-1]
		}
	}
	return paths[m-2][n-1] + paths[m-1][n-2]
}
