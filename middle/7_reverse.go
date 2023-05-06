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

func reverse(x int) int {
	strx := strconv.Itoa(int(math.Abs(float64(x))))

	firstNumFlag := false
	result := ""
	for i := len(strx) - 1; i >= 0; i-- {
		item, _ := strconv.Atoi(string(strx[i]))
		if item == 0 && !firstNumFlag {
			firstNumFlag = true
			continue
		}
		firstNumFlag = true
		result = result + string(strx[i])
	}
	r, _ := strconv.Atoi(result)
	if x < 0 {
		r = -r
	}
	if r > math.MaxInt32 || r < -math.MaxInt32 {
		return 0
	}
	return r
}

func main() {
	fmt.Println(reverse(123))
}
