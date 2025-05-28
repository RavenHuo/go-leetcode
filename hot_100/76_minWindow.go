package main

import (
	"fmt"
	"strings"
)

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	if len(s) == len(t) {
		if s == t {
			return s
		}
	}
	strMap := make(map[string]int)
	for _, s := range []byte(t) {
		strMap[string(s)] += 1
	}
	res := ""
	for i := 0; i < len(s); i++ {
		if _, ok := strMap[string(s[i])]; !ok {
			continue
		}
		temp := make(map[string]int)
		left := i + 1
		right := len(s) - 1
		temp[string(s[i])] += 1
		isFull := checkIsFull(temp, strMap)
		if isFull {
			if res == "" {
				res = s[i:left]
			} else {
				if len(res) > len(s[i:left]) {
					res = s[i:left]
				}
			}
		}
		for left <= right {
			if _, ok := strMap[string(s[left])]; ok {
				temp[string(s[left])] += 1
				isFull = checkIsFull(temp, strMap)
				if isFull {
					if res == "" {
						res = s[i : left+1]
					} else {
						if len(res) > len(s[i:left+1]) {
							res = s[i : left+1]
						}
					}
					break
				}
			}
			if !strings.Contains(t, string(s[right])) {
				right--
			}
			left++
		}
	}
	return res
}

func checkIsFull(m map[string]int, strMap map[string]int) bool {
	for s, mc := range strMap {
		c, ok := m[s]
		if !ok {
			return false
		}
		if c < mc {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(minWindow("aaaaaaaaaaaabbbbbcdd", "abcdd"))
}
