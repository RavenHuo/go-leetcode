package main

func searchMatrix2(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] > target {
			continue
		}
		left := 0
		right := len(matrix[0]) - 1
		for left <= right {
			midIndex := (left + right) / 2
			mid := matrix[i][midIndex]
			if target > mid {
				left = midIndex + 1
			} else if target == mid {
				return true
			} else if target < mid {
				right = midIndex - 1
			}
		}
	}
	return false
}
