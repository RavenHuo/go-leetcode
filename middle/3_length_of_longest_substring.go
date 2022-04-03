/**
 * @Author raven
 * @Description
 * @Date 2022/3/20
 **/
package main

// leetCode 第三题
//给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	// 双游标，i是左边的游标，j是右边的游标
	i := 0
	j := 1
	result := 0
	for {
		if j > len(s) {
			break
		}
		sub := s[i:j]
		if checkIfStringRepeat(sub) {
			i++
			if i >= j {
				j++
			}
		} else {
			if (j - i) > result {
				result = j - i
			}
			j++
		}

	}
	return result
}

func checkIfStringRepeat(str string) bool {
	if len(str) == 0 {
		return false
	}
	strMap := make(map[string]int, 0)
	for i := 0; i < len(str); i++ {
		if _, ok := strMap[str[i:i+1]]; ok {
			return true
		}
		strMap[str[i:i+1]] = 0
	}
	return false
}
