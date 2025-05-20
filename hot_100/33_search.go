package main

import "fmt"

func search(nums []int, target int) int {
	i := 0
	j := len(nums) - 1
	for i <= j {
		if nums[i] == target {
			return i
		}
		if nums[j] == target {
			return j
		}
		if nums[i] == 0 && nums[j] == 0 {
			return -1
		}
		if nums[i] > target && nums[j] < target {
			return -1
		}
		if nums[i] < target {
			i++
		}
		if nums[j] > target {
			j--
		}
	}
	return -1
}
func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
}
