/**
 * @Author raven
 * @Description
 * @Date 2022/6/13
 **/
package main

import (
	"fmt"
)

// 给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。
//
//数组中的每个元素代表你在该位置可以跳跃的最大长度。
//
//判断你是否能够到达最后一个下标
func canJump(nums []int) bool {
	i, rightBound, n := 0, nums[0], len(nums)-1
	for ; i <= rightBound && i <= n; i++ {
		rightBound = max(rightBound, i+nums[i])
	}
	return rightBound >= n
}

func max(x int, y int) int {
	if x < y {
		return y
	}
	return x
}

func main() {
	fmt.Println(canJump([]int{1, 1, 0, 1}))
}
