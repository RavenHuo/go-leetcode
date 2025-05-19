package main

// 乘积最大子数组
func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxDp := make([]int, len(nums))
	maxDp[0] = nums[0]
	minDp := make([]int, len(nums))
	minDp[0] = nums[0]
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		n := nums[i]

		maxDp[i] = maxInt(maxDp[i-1]*n, maxInt(minDp[i-1]*n, n))
		minDp[i] = minInt(maxDp[i-1]*n, minInt(minDp[i-1]*n, n))

		if max < maxDp[i] {
			max = maxDp[i]
		}
		if max < minDp[i] {
			max = minDp[i]
		}
	}
	return max
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}
