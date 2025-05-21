package main

import (
	"fmt"
	"strconv"
	"strings"
)

func decodeString(s string) string {
	path := ""
	res := ""
	temp := ""
	for i := 0; i < len(s); i++ {
		if string(s[i]) != "]" {
			path = path + string(s[i])
		} else {
			idx := 0
			for j := len(path) - 1; j > 0; j-- {
				if string(path[j]) == "[" {
					idx = j
					break
				}
			}
			chars := path[idx+1:]
			path = path[:idx]
			idx = 0
			for j := len(path) - 1; j >= 0; j-- {
				_, err := strconv.Atoi(string(path[j]))
				if err != nil {
					idx = j + 1
					break
				}
			}
			num, _ := strconv.Atoi(path[idx:])
			path = path[:idx]
			curTemp := chars + temp
			t := ""
			for j := 0; j < num; j++ {
				t += curTemp
			}
			temp = t
			if !strings.Contains(path, "[") {
				res = res + path
				res = res + temp
				path = ""
				temp = ""
			}
		}
	}
	if len(path) != 0 {
		res = res + path
	}
	return res
}

func main() {
	fmt.Println(decodeString("3[z]2[2[y]pq4[2[jk]e1[f]]]ef"))
}
