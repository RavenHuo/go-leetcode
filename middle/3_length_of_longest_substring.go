/**
 * @Author raven
 * @Description
 * @Date 2022/3/20
 **/
package main

import "fmt"

// leetCode 第三题
// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 || len(s) == 1 {
		return len(s)
	}
	i := 0
	j := i + 1
	result := 0
	for {
		if j > len(s) {
			break
		}
		subStr := s[i:j]
		// 有重复
		if checkRepeatStr(subStr) {
			if i < j {
				i++
			} else {
				j++
			}
		} else {
			if j-i > result {
				result = j - i
			}
			j++
		}
	}
	return result
}

func checkRepeatStr(s string) bool {

	m := make(map[string]struct{})
	for i := 0; i < len(s); i++ {
		item := s[i : i+1]
		if _, ok := m[item]; ok {
			return true
		} else {
			m[item] = struct{}{}
		}
	}
	return false
}

func main() {
	fmt.Println(lengthOfLongestSubstring("au"))
}
