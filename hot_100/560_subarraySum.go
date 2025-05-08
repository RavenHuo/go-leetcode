package main

import "fmt"

func subarraySum(nums []int, k int) int {
	c := 0
	if len(nums) == 0 {
		return c
	}
	if len(nums) == 1 && nums[0] == k {
		return 1
	}
	i := 1
	j := 0
	for i < len(nums) && j < len(nums) {
		subn := 0
		for n := j; n <= i; n++ {
			subn += nums[n]
		}
		if subn < k {
			if nums[i] > 0 {
				i++
			} else {
				j++
			}
		} else if subn == k {
			j++
			i++
			c += 1
		} else if subn > k {
			j++
		}
		if i < j {
			i++
		}
	}
	return c
}

func main() {
	fmt.Println(subarraySum([]int{-1, -1, 1}, 0))
}
