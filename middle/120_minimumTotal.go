package main

import "fmt"

// https://leetcode.cn/problems/triangle/description/?envType=study-plan-v2&envId=dynamic-programming
// 120. 三角形最小路径和
//中等
//相关标签
//相关企业
//给定一个三角形 triangle ，找出自顶向下的最小路径和。
//
//每一步只能移动到下一行中相邻的结点上。相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。也就是说，如果正位于当前行的下标 i ，那么下一步可以移动到下一行的下标 i 或 i + 1 。
//
//
//
//示例 1：
//
//输入：triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
//输出：11
//解释：如下面简图所示：
//   2
//  3 4
// 6 5 7
//4 1 8 3
//自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。
//示例 2：
//
//输入：triangle = [[-10]]
//输出：-10
//
//
//提示：
//
//1 <= triangle.length <= 200
//triangle[0].length == 1
//triangle[i].length == triangle[i - 1].length + 1
//-104 <= triangle[i][j] <= 104
//
//
//进阶：
//
//你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题吗？

// 每一个位置都是上一行相邻位置的最小值，加上当前位置的值
// 相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。也就是说，如果正位于当前行的下标 i ，那么下一步可以移动到下一行的下标 i 或 i + 1 。
// 反过来来说 当前位置的最小值是 上一层 i 与 i-1 的最小值
// 因为都是必须从上一层选一个位置的值，然后加到当前位置的值上面，所以不用担心溢出
func minimumTotal(triangle [][]int) int {
	m := len(triangle)
	if m == 0 {
		return 0
	}
	result := make([][]int, m)
	for i, item := range triangle {
		result[i] = make([]int, len(item))
	}
	result[0][0] = triangle[0][0]

	for i := 1; i < m; i++ {
		for j := 0; j < len(triangle[i]); j++ {
			// 当前的值
			result[i][j] = triangle[i][j]

			right := 0
			left := 0
			// 注意j的边界
			if j-1 >= 0 && j-1 <= len(result[i-1])-1 {
				left = result[i-1][j-1]
			}
			if j <= len(result[i-1])-1 {
				right = result[i-1][j]
			}
			// 最左边的边界
			if j-1 < 0 {
				result[i][j] += right
				continue
			}
			// 最右边的边界
			if j > len(result[i-1])-1 {
				result[i][j] += left
				continue
			}

			// 上一层 i 与 i-1 的最小值
			if right < left {
				result[i][j] += right
			} else {
				result[i][j] += left
			}
		}
	}
	// 最后一行的最小值就是最终的结果
	lastResult := result[m-1]
	res := lastResult[0]
	for _, item := range lastResult {
		if res > item {
			res = item
		}
	}
	return res
}

func main() {
	arr := [][]int{{7}, {-5, 9}, {6, 5, 2}, {-8, -2, -7, 3}, {-2, 6, -6, -1, 4}}
	fmt.Println(minimumTotal(arr))
}
