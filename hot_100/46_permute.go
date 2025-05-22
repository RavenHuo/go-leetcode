package main

import "fmt"

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	for _, n := range nums {
		if len(res) == 0 {
			res = append(res, []int{n})
		} else {
			temp := make([][]int, 0)
			for _, l := range res {
				temp = append(temp, insertIntoList(l, n)...)
			}
			res = temp
		}
	}
	return res
}
func insertIntoList(l []int, n int) [][]int {
	res := make([][]int, 0)
	res = append(res, append([]int{n}, l...))
	i := 1
	for i < len(l) {
		newL := make([]int, 0)
		newL = append(newL, l[:i]...)
		newL = append(newL, n)
		newL = append(newL, l[i:]...)
		res = append(res, newL)
		i++
	}
	res = append(res, append(l, n))
	return res
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}
