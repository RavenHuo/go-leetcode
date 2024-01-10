package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/3sum/
// 给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请
//
//你返回所有和为 0 且不重复的三元组。
//
//注意：答案中不可以包含重复的三元组。
//
//
//
//
//
//示例 1：
//
//输入：nums = [-1,0,1,2,-1,-4]
//输出：[[-1,-1,2],[-1,0,1]]
//解释：
//nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
//nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
//nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
//不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
//注意，输出的顺序和三元组的顺序并不重要。
//示例 2：
//
//输入：nums = [0,1,1]
//输出：[]
//解释：唯一可能的三元组和不为 0 。
//示例 3：
//
//输入：nums = [0,0,0]
//输出：[[0,0,0]]
//解释：唯一可能的三元组和为 0 。

// 1、对数组进行排序。
// 2、遍历排序后数组：
//若 nums[i]>0nums[i]>0nums[i]>0：因为已经排序好，所以后面不可能有三个数加和等于 000，直接返回结果。
//对于重复元素：跳过，避免出现重复解
//令左指针 L=i+1L=i+1L=i+1，右指针 R=n−1R=n-1R=n−1，当 L<RL<RL<R 时，执行循环：
//当 nums[i]+nums[L]+nums[R]==0nums[i]+nums[L]+nums[R]==0nums[i]+nums[L]+nums[R]==0，执行循环，判断左界和右界是否和下一位置重复，去除重复解。并同时将 L,RL,RL,R 移到下一位置，寻找新的解
//若和大于 000，说明 nums[R]nums[R]nums[R] 太大，RRR 左移
//若和小于 000，说明 nums[L]nums[L]nums[L] 太小，LLL 右移

func threeSum(nums []int) [][]int {
	// 1、先倒序 排序
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	result := make([][]int, 0)
	// 去重的map
	resultMap := make(map[string]struct{})
	for i := 0; i < len(nums); i++ {
		// 假如是最左边的key都大于0，那后面也都大于0
		if nums[i] > 0 {
			break
		}
		// 答案中不可以包含重复的三元组
		// 这个很有必要，因为已经排序过了，所以有可能存在i 跟i-1 值相同的情况，会出现重复接的情况
		// 若没有这一步 [-1,0,1,2,-1,-4]，会得到 [[-1 -1 2] [-1 0 1] [-1 0 1]] 这个解，但是这个是重复的解
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 双指针
		left := i + 1
		right := len(nums) - 1
		for left < right {

			// 左右指针相加
			sum := nums[left] + nums[right]
			if nums[i]+sum == 0 { // =0还是要移动指针
				// 去重的key
				s := toString(nums[i], nums[left], nums[right])
				// 左指针向右移动的时候，nums[left+1] == nums[left]的时候，需要去重
				if _, ok := resultMap[s]; !ok {
					result = append(result, []int{nums[i], nums[left], nums[right]})
					resultMap[s] = struct{}{}
				}

				left++
				right--
			} else if nums[i]+sum < 0 { // sum 小于零，那就是左指针要左移
				left++
			} else { // 右移
				right--
			}
		}
	}
	return result
}

func toString(ints ...int) string {
	strArr := make([]string, 0)

	for _, item := range ints {
		strArr = append(strArr, strconv.Itoa(item))
	}
	return strings.Join(strArr, "-")
}

func main() {
	arr := []int{-2, 0, 0, 2, 2}
	fmt.Println(threeSum(arr))
}
