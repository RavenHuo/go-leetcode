package main

import (
	"fmt"
	"sort"
)

// https://leetcode.cn/problems/4sum/
// 18. 四数之和
//中等
//相关标签
//相关企业
//给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。请你找出并返回满足下述全部条件且不重复的四元组 [nums[a], nums[b], nums[c], nums[d]] （若两个四元组元素一一对应，则认为两个四元组重复）：
//
//0 <= a, b, c, d < n
//a、b、c 和 d 互不相同
//nums[a] + nums[b] + nums[c] + nums[d] == target
//你可以按 任意顺序 返回答案 。
//
//
//
//示例 1：
//
//输入：nums = [1,0,-1,0,-2,2], target = 0
//输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
//示例 2：
//
//输入：nums = [2,2,2,2,2], target = 8
//输出：[[2,2,2,2]]
//

// 遍历+双指针，降低复杂度
func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return make([][]int, 0)
	}
	// 1、先倒序 排序
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	result := make([][]int, 0)
	// i,j,k,l
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			// j>i+1是避免了j跟i相同的情况，因为上面两行i==i-1已经判断过了
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			k := j + 1
			l := len(nums) - 1
			// 双指针
			for k <= l-1 {
				s := nums[i] + nums[j] + nums[k] + nums[l]
				if s == target {
					result = append(result, []int{nums[i], nums[j], nums[k], nums[l]})
					// 防止连续重复，所以需要for
					for k < len(nums)-1 && nums[k] == nums[k+1] {
						k++
					}
					// 防止连续重复，所以需要for
					for l > 0 && nums[l] == nums[l-1] {
						l--
					}
					k++
					l--
				} else if s < target {
					k++
				} else {
					l--
				}
			}
		}
	}
	return result
}

func main() {
	arr := []int{4, -9, -2, -2, -7, 9, 9, 5, 10, -10, 4, 5, 2, -4, -2, 4, -9, 5}
	result := fourSum(arr, -13)

	for _, item := range result {
		fmt.Println(item)
	}
}
