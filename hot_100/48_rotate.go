package main

func rotate2(matrix [][]int) {
	// 上下翻转+对折反转
	for i := 0; i < len(matrix)/2; i++ {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[i][j], matrix[len(matrix)-1-i][j] = matrix[len(matrix)-1-i][j], matrix[i][j]
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
