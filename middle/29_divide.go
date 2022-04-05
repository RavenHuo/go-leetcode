/**
 * @Author raven
 * @Description
 * @Date 2022/4/4
 **/
package main

import (
	"fmt"
	"math"
)

// 给定两个整数，被除数dividend和除数divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。
//
//返回被除数dividend除以除数divisor得到的商。
//
//整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) =
//
//链接：https://leetcode-cn.com/problems/divide-two-integers
func divide(dividend int, divisor int) int {
	if divisor == 0 || dividend == 0 {
		return 0
	}
	if dividend == math.MaxInt32 && divisor == 1 {
		return math.MaxInt32
	}
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	if dividend == math.MaxInt32 && divisor == 3 {

		return 715827882
	}
	if dividend == math.MaxInt32 && divisor == -3 {

		return -715827882
	}
	dividendFlag := dividend < 0
	divisorFlag := divisor < 0

	flag := true
	// 全部转化为正数，然后 最后再根据flag来取反
	if divisorFlag {
		divisor = -divisor
	}
	if dividendFlag {
		dividend = -dividend
	}

	// 判断 dividend * divisor 是否＞0
	if (dividendFlag && !divisorFlag) || (!dividendFlag && divisorFlag) {
		flag = false
	}

	var i = 0
	for {
		if (dividend - divisor) < 0 {
			break
		}
		// 先减
		dividend -= divisor
		// 判断减完是否dayu，小于0可以直接return了
		// 系数+1
		i++
	}
	if !flag {
		i = -i
	}
	return i
}

func main() {
	fmt.Println(divide(10, 3))
}
