/**
 * @Author raven
 * @Description
 * @Date 2022/5/15
 **/
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 10)
	ch <- 1
	i := <-ch
	println(i)
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	defer func() {
		fmt.Println("hello")
	}()
	go testDone(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("done ", ctx.Err())
	}

}
func testDone(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("done ", ctx.Err())
	case <-time.After(time.Duration(500) & time.Microsecond):
		fmt.Println("not done time out")
	}
}
