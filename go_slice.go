/**
 * @Author raven
 * @Description
 * @Date 2022/5/4
 **/
package main

import (
	"fmt"
)

func main() {
	//// make返回的是类型
	//slices := make([]int,2,2)
	//fmt.Printf("len :%d \n",len(slices))
	//fmt.Printf("len :%d \n",cap(slices))
	//
	//// new 返回的是指针
	//newSlices := new([]int)
	//*newSlices = append(*newSlices, 1)
	//
	//fmt.Printf("newSlices len :%d \n",len(*newSlices))
	//fmt.Printf("newSlices len :%d \n",cap(*newSlices))
	//
	//aSlice := make([]int,0,2)
	//
	//bSlice := append(aSlice, 1)
	//fmt.Printf("aSlice :%p\n", &aSlice)
	//fmt.Printf("bSlice :%p\n", &bSlice)
	//
	//arr := []int{1, 2, 3}
	//newArr := []*int{}
	//for _, v := range arr {
	//	newArr = append(newArr, &v)
	//	fmt.Println(&v)
	//	fmt.Println(v)
	//
	//}

	s1 := []int{1, 2}
	s2 := s1
	s2 = append(s2, 3)
	a(s1)
	fmt.Printf("s1:%p  len:%d \n", s1, len(s1))
	a(s2)
	fmt.Printf("s2:%p  len:%d \n", s2, len(s2))
	fmt.Println(s1, s2)
}

func a(s []int) {
	s = append(s, 0)
	fmt.Printf("%p len：%d \n", s, len(s))
	fmt.Println(s)
	for i := range s {
		s[i]++
	}
}
