package main

import (
	"fmt"
	"strings"
)

// https://leetcode.cn/problems/longest-common-prefix/
// 输入：strs = ["flower","flow","flight"]
// 输出："fl"
func longestCommonPrefix(strs []string) string {
	result := ""
	if len(strs) == 0 {
		return result
	}
	for i := 1; i <= len(strs[0]); i++ {
		result = strs[0][:i]
		for j := 1; j < len(strs); j++ {
			if !strings.HasPrefix(strs[j], result) {
				return result
			}
		}
	}
	return result
}

func main() {
	strs := []string{"a"}
	fmt.Println(longestCommonPrefix(strs))
}
