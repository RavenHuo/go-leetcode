package main

func isPalindromeStr(s string) bool {
	if len(s) == 0 {
		return false
	}
	l := 0
	r := len(s) - 1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
