package main

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 || len(intervals) == 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	cur := make([]int, 0)
	resultList := make([][]int, 0)
	i := 0
	for i < len(intervals) {
		if len(cur) == 0 {
			cur = intervals[i]
		} else {
			l := intervals[i][0]
			r := intervals[i][1]
			if cur[0] <= r && cur[1] >= l {
				if l < cur[0] {
					cur[0] = l
				}
				if cur[1] < r {
					cur[1] = r
				}
			} else {
				resultList = append(resultList, cur)
				cur = intervals[i]
			}
		}
		i++
	}
	resultList = append(resultList, cur)
	return resultList
}
