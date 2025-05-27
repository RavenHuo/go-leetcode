package main

import "sort"

// 本质上就是固定一位然后双指针，每次移动看当前值和0的关系，决定下一次移动左边还是移动右边。
func threeSum(nums []int) [][]int {
	// 先排序
	sort.Ints(nums)
	result := make([][]int, 0)
	for i, num := range nums {
		left := i + 1
		// 去重
		if i-1 >= 0 && num == nums[i-1] {
			continue
		}
		right := len(nums) - 1
		for left < right {
			if num+nums[left]+nums[right] == 0 {
				result = append(result, []int{num, nums[left], nums[right]})
				left++
				// 去重
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				continue
			}
			if num+nums[left]+nums[right] < 0 {
				left++
				continue
			}
			if num+nums[left]+nums[right] > 0 {
				right--
				continue
			}
		}
	}
	return result
}
