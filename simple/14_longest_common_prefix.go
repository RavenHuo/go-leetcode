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

	for pIndex := 1; pIndex <= len(strs[0]); pIndex++ {
		prefix := strs[0][:pIndex]
		for i := 0; i < len(strs); i++ {
			if !strings.HasPrefix(strs[i], prefix) {
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
