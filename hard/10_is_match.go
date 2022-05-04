/**
 * @Author raven
 * @Description
 * @Date 2022/4/5
 **/
package main

import (
	"fmt"
)

// 给你一个字符串s和一个字符规律p，请你来实现一个支持 '.'和'*'的正则表达式匹配。
//
//'.' 匹配任意单个字符
//'*' 匹配零个或多个前面的那一个元素
//所谓匹配，是要涵盖整个字符串s的，而不是部分字符串。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/regular-expression-matching
func isMatch(s string, p string) bool {
	if s == p {
		return true
	}
	sLen := len(s)
	pLen := len(p)
	i := 0
	j := 0
	starFlag := false
	for {
		if i >= sLen || j >= pLen {
			break
		}
		sChar := s[i : i+1]
		pChar := p[j : j+1]
		if pChar == "." {
			pChar = sChar
		}
		if pChar == "*" {
			if j-1 < 0 {
				return false
			}
			pChar = p[j-1 : j]
			if pChar == "." {
				pChar = sChar
			}
			starFlag = true
		} else {
			starFlag = false
		}
		// *替换之后 也不相等的时候，就要j++ ，因为aab 跟a*b 是匹配的
		// aab 跟c*a*b 也是匹配的，因为*可以匹配0个或多个
		if pChar != sChar {
			if starFlag {
				j++
				continue
			} else {
				if j+1 >= pLen {
					return false
				} else if p[j+1:j+2] == "*" {
					// 不相等的情况下，后一个是* 那就代表匹配0个
					j++
					continue
				} else {
					return false
				}
			}
		} else {
			if p[j:j+1] == "*" && j < pLen-1 {
				// a*a == a 处理这种情况
				// 相等的情况下，
				for {
					if pChar != sChar {
						break
					}

				}
			}
		}

	}
	if i == sLen {
		if j == pLen {
			return true
		} else {
			if j == pLen-1 && p[j:j+1] == "*" && (p[j-1:j] == s[i-1:i] || p[j-1:j] == ".") {
				return true
			} else if j == pLen-2 && p[j:j+1] == "*" && (p[j+1:j+2] == s[i-1:i]) {
				return true
			} else {
				return false
			}
		}
	} else {
		return false
	}
}

// "mississippi"
// "mis*is*p*."
// "aaa"
//"ab*a*c*a"
func main() {
	fmt.Println(isMatch("aaa", "ab*a*c*a"))
}
