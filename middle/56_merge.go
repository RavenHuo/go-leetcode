package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	result := make([][]int, 0, len(intervals))
	if len(intervals) == 0 {
		return result
	}
	if len(intervals) == 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	preStart := intervals[0][0]
	preEnd := intervals[0][1]

	result = append(result, []int{preStart, preEnd})
	preIndex := 0

	for i := 1; i < len(intervals); i++ {

		start := intervals[i][0]
		end := intervals[i][1]
		if start >= preStart && end <= preEnd {
			continue
		} else if end >= preStart && preEnd >= start { // 交集
			result[preIndex] = []int{preStart, end}
			preEnd = end // 更新
		} else {
			result = append(result, []int{start, end})
			preStart = start
			preEnd = end
			preIndex++
		}
	}

	return result
}

func main() {
	arrList := merge([][]int{{0, 4}, {3, 5}})
	for _, arr := range arrList {
		fmt.Println(arr)
	}
}
