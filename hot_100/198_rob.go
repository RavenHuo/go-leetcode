package main

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return maxInt(nums[0], nums[1])
	}

	res := make([]int, len(nums))
	res[0] = nums[0]
	res[1] = nums[1]
	m := maxInt(res[0], res[1])
	n := 2
	for n < len(nums) {
		if n-3 < 0 {
			res[n] = nums[n-2] + nums[n]
		} else {
			res[n] = maxInt(res[n-2], res[n-3]) + nums[n]
		}
		m = maxInt(res[n], m)
		n++
	}
	return m
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
