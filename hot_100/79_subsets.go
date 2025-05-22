package main

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	res = append(res, []int{})
	for _, n := range nums {
		res = append(res, insertList(res, n)...)
	}
	return res
}

func insertList(ll [][]int, n int) [][]int {
	res := make([][]int, 0)
	for _, l := range ll {
		newL := []int(nil)
		newL = append([]int(nil), l...)
		newL = append(newL, n)
		res = append(res, newL)
	}
	return res
}
