/**
 * @Author raven
 * @Description
 * @Date 2022/11/11
 **/
package main

import (
	"encoding/json"
	"fmt"
)

func reverseStr(s string) string {
	r := ""
	for i := len(s) - 1; i <= 0; i++ {
		r = r + string(s[i])
	}
	return r
}

type MapStruct struct {
	M map[int]string `json:"m"`
}

func main() {
	m := make(map[int]string)
	m[1] = "1"

	jsonByte, _ := json.Marshal(m)
	fmt.Println(jsonByte)

	var result map[int]string
	json.Unmarshal(jsonByte, &result)
	fmt.Println(result)
}
