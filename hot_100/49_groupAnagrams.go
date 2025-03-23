package main

import "sort"

func groupAnagrams(strs []string) [][]string {
	// 保存下标
	kMap := make(map[string]int)
	result := make([][]string, 0)
	for _, str := range strs {
		sortStr := getSortStr(str)
		if idx, ok := kMap[sortStr]; !ok {
			kMap[sortStr] = len(result)
			result = append(result, []string{str})
		} else {
			result[idx] = append(result[idx], str)
		}
	}
	return result
}

func getSortStr(s string) string {
	if s == "" {
		return s
	}
	ch := []byte(s)
	sort.Slice(ch, func(i, j int) bool {
		return ch[i] < ch[j]
	})
	return string(ch)
}

func main() {
	groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
}
