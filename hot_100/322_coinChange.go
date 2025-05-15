package main

func coinChange(coins []int, amount int) int {
	// 存结果
	nums := make([]int, amount+1)
	i := 1
	for i <= amount {
		curInt := -1
		for _, n := range coins {
			if n == i {
				curInt = 1
				break
			}
			if i-n >= 0 && nums[i-n] != -1 { // 找下标
				if curInt == -1 {
					curInt = nums[i-n] + 1
				} else {
					curInt = minInt(curInt, nums[i-n]+1)
				}
			}
		}
		nums[i] = curInt
		i++
	}
	return nums[amount]
}

func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	coinChange([]int{2}, 3)
}
