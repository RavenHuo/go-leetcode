package main

// 找到最便宜买点，那后面股价最高的点就是利润最高的点
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	minPrice, profit := prices[0], 0
	for _, p := range prices {
		if p < minPrice {
			minPrice = p
		}
		curProfit := p - minPrice
		if curProfit > profit {
			profit = curProfit
		}
	}
	return profit
}
