package main

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	res := make([]int, 0)
	top := 0
	down := len(matrix) - 1
	left := 0
	right := len(matrix[0]) - 1
	for len(res) < len(matrix)*len(matrix[0]) {
		if top <= down {
			for i := left; i <= right; i++ {
				res = append(res, matrix[top][i])
			}
			// 上边下移
			top += 1
		}
		if left <= right {
			for i := top; i <= down; i++ {
				res = append(res, matrix[i][right])
			}
			// 右边右移动
			right -= 1
		}
		if top <= down {
			for i := right; i >= left; i-- {
				res = append(res, matrix[down][i])
			}
			down -= 1
		}
		if left <= right {
			for i := down; i >= top; i-- {
				res = append(res, matrix[i][left])
			}
			left += 1
		}
	}
	return res
}
