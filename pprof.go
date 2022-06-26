/**
 * @Author raven
 * @Description
 * @Date 2022/6/3
 **/
package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/pprof"
	"strconv"
	"time"
)

func main() {
	profile()
	// 该方法需要在listen方法前使用，作用使先注册路由
	http.HandleFunc("/debug/pprof/heap", pprof.Index)
	//http.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	//http.HandleFunc("/debug/pprof/profile", pprof.Profile)
	//http.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	//http.HandleFunc("/debug/pprof/trace", pprof.Trace)
	err := http.ListenAndServe("localhost:8080", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}

func profile() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				select {
				case <-time.After(time.Second):
					fmt.Println("hello :" + strconv.Itoa(i))
				}
			}
		}(i)
	}
}
