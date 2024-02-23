package main

import (
	"fmt"
	"strings"
)

// https://leetcode.cn/problems/longest-common-prefix/
// 输入：strs = ["flower","flow","flight"]
// 输出："fl"
func longestCommonPrefix(strs []string) string {
	commonPrefix := ""
	if len(strs) == 0 {
		return commonPrefix
	}
	for i := 1; i <= len(strs[0]); i++ {
		prefix := strs[0][:i]
		for j := 1; j < len(strs); j++ {
			if !strings.HasPrefix(strs[j], prefix) {
				return commonPrefix
			}
		}
		commonPrefix = prefix
	}
	return commonPrefix
}

func main() {
	strs := []string{"a"}
	fmt.Println(longestCommonPrefix(strs))
}
