/**
 * @Author raven
 * @Description
 * @Date 2022/11/11
 **/
package main

import (
	"fmt"
	"math"
	"strconv"
)

// 给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。
//
// 如果反转后整数超过 32 位的有符号整数的范围 [−231,  231 − 1] ，就返回 0。
func reverse(x int) int {
	str := strconv.Itoa(int(math.Abs(float64(x))))
	reverseStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		zero, _ := strconv.Atoi(string(str[i]))
		if zero == 0 {
			continue
		}
		reverseStr += string(str[i])
	}
	r, _ := strconv.Atoi(reverseStr)
	if x < 0 {
		r = -r
	}
	if r > math.MaxInt32 || r < math.MinInt32 {
		return 0
	}
	return r
}

func main() {
	fmt.Println(reverse(1230))
}
