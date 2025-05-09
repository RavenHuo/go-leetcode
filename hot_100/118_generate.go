package main

import "fmt"

func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	for i := 1; i <= numRows; i++ {
		arr := make([]int, i)
		if i == 1 || i == 2 {
			arr[0] = 1
			arr[i-1] = 1
		} else {
			arr[0] = 1
			for n := 1; n < i-1; n++ {
				arr[n] = result[i-2][n-1] + result[i-2][n]
			}
			arr[i-1] = 1
		}
		result[i-1] = arr
	}
	return result
}

func main() {
	fmt.Println(generate(5))
}
