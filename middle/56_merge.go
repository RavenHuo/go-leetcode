package main

import (
	"fmt"
	"sort"
)

// [a,b] [c,d]
// c>=a && d<=b 包含关系
// d>=a && b>=c 交集
func merge(intervals [][]int) [][]int {
	result := make([][]int, 0)
	if len(intervals) == 0 {
		return result
	}
	if len(intervals) == 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	firstArr := intervals[0]
	preStart := firstArr[0]
	preEnd := firstArr[1]
	result = append(result, []int{preStart, preEnd})
	preIndex := 0
	for i := 1; i < len(intervals); i++ {
		start := intervals[i][0]
		end := intervals[i][1]
		if preStart <= start && end <= preEnd { // 包含
			continue
		} else if preStart <= end && preEnd >= start { // 交集
			result[preIndex][1] = end
			preEnd = end
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
