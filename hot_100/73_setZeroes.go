package main

import "math"

func setZeroes(matrix [][]int) {
	for i := range len(matrix) {
		for j := range len(matrix[i]) {
			if matrix[i][j] == 0 {
				for k := range len(matrix[i]) {
					if k > j && matrix[i][k] != 0 {
						matrix[i][k] = math.MaxInt
					} else {
						matrix[i][k] = 0
					}
				}
				for k := range len(matrix) {
					if k > i && matrix[k][j] != 0 {
						matrix[k][j] = math.MaxInt
					} else {
						matrix[k][j] = 0
					}
				}
			} else if matrix[i][j] == math.MaxInt {
				matrix[i][j] = 0
			}
		}
	}
}
