/**
 * @Author raven
 * @Description
 * @Date 2022/9/29
 **/
package main

func longestPalindrome(s string) string {
	res, dp := "", make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	// i从右边往左边走
	for i := len(s) - 1; i >= 0; i-- {
		// j从 i的右边 往字符串的末尾走
		for j := i; j < len(s); j++ {
			// 滑动窗口公式
			dp[i][j] = (s[i] == s[j]) && ((j-i < 3) || dp[i+1][j-1])
			if dp[i][j] && (res == "" || j-i+1 > len(res)) {
				res = s[i : j+1]
			}
		}
	}
	return res
}
