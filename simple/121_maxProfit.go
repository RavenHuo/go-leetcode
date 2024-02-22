package main

// 给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
//
// 你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
//
// 返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/description/

// 找到最小的price
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	m := 0
	minPrice := prices[0]
	for _, price := range prices {
		// 当前价格减去最小值等于利润，因为卖出要在买入之后，所以这个顺序是没问题的
		m = max(m, price-minPrice)
		minPrice = min(price, minPrice)
	}
	return m
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
