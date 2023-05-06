/**
 * @Author raven
 * @Description
 * @Date 2022/11/15
 **/
package main

import "math"

// 缺失的第一个正数 0不算正整数
func firstMissingPositive(nums []int) int {
	numsMap := make(map[int]bool, 0)
	for i := 0; i < len(nums); i++ {
		numsMap[nums[i]] = true
	}
	for i := 1; i < math.MaxInt32; i++ {
		if !numsMap[i] {
			return i
		}
	}
	return 0
}
