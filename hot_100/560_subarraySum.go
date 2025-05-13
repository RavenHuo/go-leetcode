package main

import "fmt"

// 前缀和+两数之和
func subarraySum(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 && nums[0] == k {
		return 1
	}
	c := 0
	rMap := make(map[int]int)
	rMap[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == k {
			c++
			rMap[i] = nums[i]
			continue
		}
		if rMap[i-1]+nums[i] < rMap[i-1] {
			rMap[i] = nums[i]
			continue
		}
		if rMap[i-1]+nums[i] == k {
			c++
			rMap[i] = nums[i]
		} else if rMap[i-1]+nums[i] < k {
			rMap[i] = rMap[i-1] + nums[i]
		} else if rMap[i-1]+nums[i] > k {
			rMap[i] = nums[i]
		}
	}
	return c
}

func main() {
	fmt.Println(subarraySum([]int{-1, -1, 1}, 1))
}
