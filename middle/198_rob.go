package main

import "fmt"

// https://leetcode.cn/problems/house-robber/?envType=study-plan-v2&envId=dynamic-programming
// 198. 打家劫舍
//中等
//2.9K
//相关企业
//你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
//
//给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。
//
//
//
//示例 1：
//
//输入：[1,2,3,1]
//输出：4
//解释：偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
//     偷窃到的最高金额 = 1 + 3 = 4 。
//示例 2：
//
//输入：[2,7,9,3,1]
//输出：12
//解释：偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。
//     偷窃到的最高金额 = 2 + 9 + 1 = 12 。
//
//
//提示：
//
//1 <= nums.length <= 100
//0 <= nums[i] <= 400

// 每一个位置的最优解是 前面[:n-1]数组的max + 当前位置的值
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		if gt(nums[0], nums[1]) {
			return nums[0]
		}
		return nums[1]
	}
	result := make([]int, len(nums))
	result[0] = nums[0]
	result[1] = nums[1]
	for i := 2; i < n; i++ {
		// 前面的最大值+当前值等于目前位置的最优解
		m := maxArr(result[:i-1])
		result[i] = m + nums[i]
	}

	return maxArr(result)
}

func gt(i, j int) bool {
	return i >= j
}

func maxArr(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	maxResult := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > maxResult {
			maxResult = arr[i]
		}
	}
	return maxResult
}

func main() {
	arr := []int{1, 2, 3, 1}
	fmt.Println(rob(arr))
}
