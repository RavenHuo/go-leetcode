package main

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 || len(s) == 1 {
		return len(s)
	}
	m := make(map[byte]struct{})
	l := 0
	chars := []byte(s)
	i := 0
	j := 0
	for i > len(s)-1 {
		char := chars[i]
		_, ok := m[char]
		if !ok {
			m[char] = struct{}{}
			if l < len(m) {
				l = len(m)
			}
			i++
		} else {
			// 有重复
			for i > j {
				if chars[j] == char {
					j++
					break
				}
				delete(m, chars[j])
				j++
			}
		}
	}
	return l
}
