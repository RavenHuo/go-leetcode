package main

// https://leetcode.cn/problems/unique-paths/?envType=study-plan-v2&envId=dynamic-programming
// 62. 不同路径
//中等
//相关标签
//相关企业
//一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
//
//机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
//
//问总共有多少条不同的路径？
//
//
//
//示例 1：
//
//
//输入：m = 3, n = 7
//输出：28
//示例 2：
//
//输入：m = 3, n = 2
//输出：3
//解释：
//从左上角开始，总共有 3 条路径可以到达右下角。
//1. 向右 -> 向下 -> 向下
//2. 向下 -> 向下 -> 向右
//3. 向下 -> 向右 -> 向下
//示例 3：
//
//输入：m = 7, n = 3
//输出：28
//示例 4：
//
//输入：m = 3, n = 3
//输出：6
//
//
//提示：
//
//1 <= m, n <= 100
//题目数据保证答案小于等于 2 * 109

// 机器人每次只能向下或者向右移动一步，那就是每一个位置的路径的总和是（当前位置的左边的位置的路径总和+当前位置的上边的位置的路径总和）
// f(m,n) = f(m-1,n)+f(m,n-1)
func uniquePaths(m int, n int) int {
	if m == 0 {
		return n
	}
	if n == 0 {
		return m
	}
	result := make([][]int, m)
	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
	}
	// 只有一格的时候是 1 只有一条路
	result[0][0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// i==0 的时候只能向右走
			if i > 0 {
				result[i][j] += result[i-1][j]
			}
			// j==0的时候只能向下走
			if j > 0 {
				result[i][j] += result[i][j-1]
			}
		}
	}
	return result[m-1][n-1]
}

func main() {
	println(uniquePaths(3, 7))
}
