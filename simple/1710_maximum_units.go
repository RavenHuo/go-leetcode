/**
 * @Author raven
 * @Description
 * @Date 2022/11/15
 **/
package main

import (
	"fmt"
	"sort"
)

// 请你将一些箱子装在 一辆卡车 上。给你一个二维数组 boxTypes ，其中 boxTypes[i] = [numberOfBoxesi, numberOfUnitsPerBoxi] ：
//
//numberOfBoxesi 是类型 i 的箱子的数量。
//numberOfUnitsPerBoxi 是类型 i 每个箱子可以装载的单元数量。
//整数 truckSize 表示卡车上可以装载 箱子 的 最大数量 。只要箱子数量不超过 truckSize ，你就可以选择任意箱子装到卡车上。
//
//返回卡车可以装载 单元 的 最大 总数。

func maximumUnits(boxTypes [][]int, truckSize int) int {
	// 质量数组
	perArr := make([]int, 0, len(boxTypes))
	//k：质量，v：数量
	perMap := make(map[int]int, len(boxTypes))
	for i := 0; i < len(boxTypes); i++ {
		if _, ok := perMap[boxTypes[i][1]]; ok {
			perMap[boxTypes[i][1]] += boxTypes[i][0]
		} else {
			perMap[boxTypes[i][1]] = boxTypes[i][0]
			perArr = append(perArr, boxTypes[i][1])
		}

	}
	//质量排序
	sort.Slice(perArr, func(i, j int) bool {
		return perArr[i] > perArr[j]
	})
	cur := 0
	result := 0
	for i := 0; i < len(perArr); i++ {
		// 数量
		for j := 0; j < perMap[perArr[i]]; j++ {
			if cur+1 > truckSize {
				// 跳出当前循环
				return result
			}
			cur++
			result += perArr[i]
		}
	}
	return result
}

func main() {

	fmt.Println(maximumUnits([][]int{{1, 3}, {5, 5}, {2, 5}, {4, 2}, {4, 1}, {3, 1}, {2, 2}, {1, 3}, {2, 5}, {3, 2}}, 35))
}
