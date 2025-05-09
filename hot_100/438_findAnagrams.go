package main

import "fmt"

// 滑动窗口，窗口大小为len(p)
func findAnagrams(s string, p string) []int {
	if s == "" {
		return nil
	}
	if len(s) < len(p) {
		return nil
	}
	// 入参的map
	pMap := make(map[byte]int)
	// 长字符串的byte数组
	sBytes := []byte(s)
	//结果集
	rMap := make(map[byte]int)
	resultList := make([]int, 0)

	for _, char := range []byte(p) {
		pMap[char] += 1
	}
	t := len(p)
	i := t - 1
	j := 0

	for n := j; n <= i; n++ {
		rMap[sBytes[n]] += 1
	}

	for i <= len(s)-1 {
		// 结果一样
		if compareMap(rMap, pMap) {
			resultList = append(resultList, j)
		}
		// 防止溢出
		if i+1 > len(s)-1 {
			break
		}

		// 先移除旧的，再加入新的
		rMap[sBytes[j]] -= 1
		i++
		j++
		rMap[sBytes[i]] += 1

	}

	return resultList
}

func compareMap(m1, m2 map[byte]int) bool {
	fmt.Printf("%+v\n", m1)
	for k, v := range m2 {
		if v != m1[k] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Print(findAnagrams("cbaebabacd", "abc"))
}
