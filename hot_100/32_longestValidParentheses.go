package main

import "fmt"

func longestValidParentheses(s string) int {
	if len(s) == 0 {
		return 0
	}
	charList := make([]string, 0)
	dp := make([]int, len(s))
	dp[0] = 0
	max := 0
	count := 0
	for i := 0; i < len(s); i++ {
		cur := string(s[i])
		if cur == ")" {
			if len(charList) > 0 {
				if charList[len(charList)-1] == "(" {
					count += 1
					charList = charList[:len(charList)-1]
					if i-2*count >= 0 {
						dp[i] = dp[i-2*count] + 2*count
					} else {
						dp[i] = 2
					}
					if dp[i] > max {
						max = dp[i]
					}
				}
			}
		} else {
			count = 0
			charList = append(charList, string(s[i]))
		}
	}

	return max
}
func main() {
	fmt.Println(longestValidParentheses("()(())"))
}
