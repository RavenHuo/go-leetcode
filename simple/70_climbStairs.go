package main

import (
	"fmt"
	"strconv"
)

// https://leetcode.cn/problems/climbing-stairs/description/?envType=study-plan-v2&envId=dynamic-programming
// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//
//每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
//
//
//
//示例 1：
//
//输入：n = 2
//输出：2
//解释：有两种方法可以爬到楼顶。
//1. 1 阶 + 1 阶
//2. 2 阶
//示例 2：
//
//输入：n = 3
//输出：3
//解释：有三种方法可以爬到楼顶。
//1. 1 阶 + 1 阶 + 1 阶
//2. 1 阶 + 2 阶
//3. 2 阶 + 1 阶
//
//
//提示：
//
//1 <= n <= 45

func main() {
	fmt.Printf(strconv.Itoa(climbStairs(1)))
}

func climbStairs(n int) int {
	if n == 1 {
		return n
	}
	if n == 2 {
		return n
	}
	climbArr := make([]int, 0)
	climbArr[0] = 1
	climbArr[1] = 2
	for i := 2; i < n; i++ {
		climbArr[i] = climbArr[i-1] + climbArr[i-2]
	}
	return climbArr[n-1]
}
