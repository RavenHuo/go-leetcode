package main

func climbStairs(n int) int {

	if n <= 2 {
		return n
	}
	nums := make([]int, n)
	nums[0] = 1
	nums[1] = 2
	i := 2
	for i < n {
		nums[i] = nums[i-1] + nums[i-2]
		i++
	}
	return nums[n-1]
}
