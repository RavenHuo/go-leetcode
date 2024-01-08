package main

// https://leetcode.cn/problems/unique-paths-ii/description/?envType=study-plan-v2&envId=dynamic-programming
// 63. 不同路径 II
//中等
//相关标签
//相关企业
//提示
//一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
//
//机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish”）。
//
//现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？
//
//网格中的障碍物和空位置分别用 1 和 0 来表示。
//
//
//
//示例 1：
//
//
//输入：obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]
//输出：2
//解释：3x3 网格的正中间有一个障碍物。
//从左上角到右下角一共有 2 条不同的路径：
//1. 向右 -> 向右 -> 向下 -> 向下
//2. 向下 -> 向下 -> 向右 -> 向右
//示例 2：
//
//
//输入：obstacleGrid = [[0,1],[0,0]]
//输出：1
//
//
//提示：
//
//m == obstacleGrid.length
//n == obstacleGrid[i].length
//1 <= m, n <= 100
//obstacleGrid[i][j] 为 0 或 1

// 每一个位置的路径总和都是 （左边的位置路径+右边的位置路径）之和
// 每一个位置的路径总和都是 （左边的位置路径+右边的位置路径）之和，跟不同的路径差不多，只是需要判断一下当前位置是不是=1，若=1就continue，当前位置的result就是0
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m == 0 {
		return 0
	}
	n := len(obstacleGrid[0])

	result := make([][]int, m)
	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
	}
	// 第一个左上角的位置就有石头，那应该是走不下去了
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	result[0][0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				continue
			}
			if i > 0 {
				result[i][j] += result[i-1][j]
			}
			if j > 0 {
				result[i][j] += result[i][j-1]
			}
		}
	}
	return result[m-1][n-1]
}
