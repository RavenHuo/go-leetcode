package main

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, n := range nums {
		need := target - n
		res, ok := numMap[need]
		if ok {
			return []int{i, res}
		}
		numMap[n] = i
	}
	return []int{}
}
