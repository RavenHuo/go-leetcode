/**
 * @Author raven
 * @Description
 * @Date 2022/5/14
 **/
package main

import (
	"fmt"
)

func main() {
	str := "111"
	fmt.Println(str[0])
	defer fmt.Println("in main")
	defer func() {
		defer func() {
			panic("panic again and again")
		}()
		panic("panic again")
	}()

	panic("panic once")
}
