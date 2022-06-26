/**
 * @Author raven
 * @Description
 * @Date 2022/5/7
 **/
package main

import (
	"fmt"
)

func main() {
	//data := make(map[int]int, 0)
	//wp := sync.WaitGroup{}
	//for i := 0; i < 100; i++ {
	//	wp.Add(1)
	//	go func(wp *sync.WaitGroup, num int) {
	//		data[num] = num
	//		defer wp.Done()
	//	}(&wp, i)
	//}
	//wp.Wait()
	//fmt.Println(len(data))
	var s *TestStruct
	fmt.Println(s == nil)    // #=> true
	fmt.Println(NilOrNot(s)) // #=> false

}

type TestStruct struct{}

func NilOrNot(v interface{}) bool {
	return v == nil
}
