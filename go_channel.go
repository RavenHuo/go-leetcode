/**
 * @Author raven
 * @Description
 * @Date 2022/5/29
 **/
package main

import "fmt"

// 最大值-最小值
//返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。
//示例 1：
//输入：[7,1,5,3,6,4]
//输出：5
//解释：在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
//     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。
//示例 2：
//输入：prices = [7,6,4,3,1]
//输出：0
//解释：在这种情况下, 没有交易完成, 所以最大利润为 0。

func maxPro(val []int) int {
	if len(val) == 0 {
		return 0
	}
	minVal := val[0]
	maxP := 0
	for _, v := range val {
		minVal = min(minVal, v)
		maxP = max(maxP, v-minVal)
	}
	return maxP
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	arr := []int{1, 2, 3, 4}
	fmt.Println(len(arr))
	fmt.Println(cap(arr))
}
