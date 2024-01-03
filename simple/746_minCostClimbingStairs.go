package main

import "fmt"

// https://leetcode.cn/problems/min-cost-climbing-stairs/?envType=study-plan-v2&envId=dynamic-programming
// 746. 使用最小花费爬楼梯
//提示
//简单
//1.4K
//相关企业
//给你一个整数数组 cost ，其中 cost[i] 是从楼梯第 i 个台阶向上爬需要支付的费用。一旦你支付此费用，即可选择向上爬一个或者两个台阶。
//
//你可以选择从下标为 0 或下标为 1 的台阶开始爬楼梯。
//
//请你计算并返回达到楼梯顶部的最低花费。
//
//
//
//示例 1：
//
//输入：cost = [10,15,20]
//输出：15
//解释：你将从下标为 1 的台阶开始。
//- 支付 15 ，向上爬两个台阶，到达楼梯顶部。
//总花费为 15 。
//示例 2：
//
//输入：cost = [1,100,1,1,1,100,1,1,100,1]
//输出：6
//解释：你将从下标为 0 的台阶开始。
//- 支付 1 ，向上爬两个台阶，到达下标为 2 的台阶。
//- 支付 1 ，向上爬两个台阶，到达下标为 4 的台阶。
//- 支付 1 ，向上爬两个台阶，到达下标为 6 的台阶。
//- 支付 1 ，向上爬一个台阶，到达下标为 7 的台阶。
//- 支付 1 ，向上爬两个台阶，到达下标为 9 的台阶。
//- 支付 1 ，向上爬一个台阶，到达楼梯顶部。
//总花费为 6 。
//
//
//提示：
//
//2 <= cost.length <= 1000
//0 <= cost[i] <= 999

// 一个台阶分两步，其实也就是每一步的最优解是前一步的最优解或者前两步的最优解

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return cost[0]
	}
	if n == 2 {
		if lt(cost[0], cost[1]) {
			return cost[0]
		}
		return cost[1]
	}
	result := make([]int, len(cost)+n)
	result[0] = cost[0]
	result[1] = cost[1]
	// n+1是因为要走到最后
	for i := 2; i < n; i++ {
		// 最优选择了i-2
		if lt(result[i-2], result[i-1]) {
			result[i] = result[i-2] + cost[i]
		} else { // 最优选择了i-1
			result[i] = result[i-1] + cost[i]
		}
	}
	// 最后也是需要比较一下 n-1跟n-2那个是最优解
	if lt(result[n-2], result[n-1]) {
		return result[n-2]
	}
	return result[n-1]
}

func lt(i, j int) bool {
	return i <= j
}

func main() {
	arr := []int{10, 15, 20}
	fmt.Println(minCostClimbingStairs(arr))
}
