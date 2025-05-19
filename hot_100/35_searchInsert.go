package main

func searchInsert(nums []int, target int) int {
	i := 0
	j := len(nums) - 1
	for i <= j {
		if target == nums[i] {
			return i
		}
		if target == nums[j] {
			return j
		}
		if target > nums[i] {
			i++
		}
		if target < nums[j] {
			j--
		}
	}
	return i
}
