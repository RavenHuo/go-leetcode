package main

import "math"

func findMin(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	left := 0
	right := math.MaxInt
	i := 0
	j := len(nums) - 1
	if nums[i] < nums[j] {
		return nums[i]
	}
	for i <= j {
		if right == left {
			return nums[i]
		}
		if nums[i] >= left {
			left = nums[i]
			i++
			continue
		}
		if nums[j] <= right {
			right = nums[j]
			j--
			continue
		}
	}
	return nums[i]
}
