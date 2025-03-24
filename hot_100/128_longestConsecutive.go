package main

import "sort"

func longestConsecutive(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}

	sort.Ints(nums)
	// 有序数组
	arr := make([]int, 0)
	arr = append(arr, nums[0])
	l := 1
	for i := 1; i < len(nums); i++ {
		n := nums[i]
		// 连续
		if n == arr[len(arr)-1]+1 { // 塞入到数组最后面
			arr = append(arr, n)
			if len(arr) > l {
				l = len(arr)
			}
		} else if n == arr[0]-1 { // 塞入到数组第一位
			newArr := make([]int, 0)
			newArr = append(newArr, n)
			newArr = append(newArr, arr...)
			arr = newArr
			if len(arr) > l {
				l = len(arr)
			}
		} else if n == arr[len(arr)-1] || n == arr[0] { // 相同跳过
			continue
		} else {
			arr = make([]int, 0)
			arr = append(arr, n)
		}
	}
	return l
}
