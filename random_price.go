/**
 * @Author raven
 * @Description
 * @Date 2022/11/21
 **/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 随机金额算法 二倍均值法
func randomPrice(amount, min, num int) {
	if amount < num {
		num = amount
	}
	// 剩下多少钱
	remain := amount - min*num
	rand.Seed(time.Now().UnixNano())
	sum := 0
	for i := 0; i < num; i++ {
		redpeck := 0
		random := rand.Intn(100)
		// 最后一个就等于剩余的钱
		if i == num-1 {
			redpeck = remain
		} else {
			redpeck = remain * random * 2 / (num - i) / 100
		}
		if remain > redpeck {
			remain = remain - redpeck
		} else {
			remain = 0
		}
		redpeckNum := min + redpeck
		sum = sum + redpeckNum
		fmt.Printf("第%d个红包金额是:%d  \n", i, redpeckNum)
	}
	fmt.Printf("红包总额是：%d", sum)
}

// 随机金额算法 二倍均值法
func randomPriceArr(amount, min, num int) []int {
	priceResult := make([]int, 0, num)
	// 剩下多少钱
	remain := amount - min*num
	rand.Seed(time.Now().UnixNano())
	sum := 0
	for i := 0; i < num; i++ {
		redpeck := 0
		random := rand.Intn(100)
		// 最后一个就等于剩余的钱
		if i == num-1 {
			redpeck = remain
		} else {
			redpeck = remain * random * 2 / (num - i) / 100
		}
		if remain > redpeck {
			remain = remain - redpeck
		} else {
			remain = 0
		}
		redpeckNum := min + redpeck
		priceResult = append(priceResult, redpeckNum)
		sum = sum + redpeckNum
		fmt.Printf("第%d个红包金额是:%d  \n", i, redpeckNum)
	}
	fmt.Printf("红包总额是：%d \n", sum)
	return priceResult
}

func main() {
	//randomPrice(100, 1, 15)
	//randomPrice(100, 1, 5)
	//randomPrice(100, 1, 5)
	//randomPrice(100, 1, 5)
	//
	//rids := []int{40109399,31218001,26863137,49852991,39553163,40344579,46751025,41217701,43889603,41880571,52495042,37148511,35525163,33894928,45278067,51511218,52605978,51100700,37816102,40160384,52476438,48833897,50722718}
	//list := make([]*RoomLimitConfig,0,len(rids))
	//for _,item:= range rids {
	//	list = append(list, &RoomLimitConfig{
	//		RoomId: item,
	//		WhiteList: []int{},
	//		LimitNum: 1000,
	//	})
	//}
	//bytes,_:=json.Marshal(list)
	//fmt.Println(string(bytes))

	intArr := make([]int, 0)
	for i := 0; i < 100; i++ {
		intArr = append(intArr, i)
	}

	for i, _ := range intArr {
		b := intArr[i]
		go func() {
			fmt.Println(b)
		}()
	}
	time.Sleep(10 * time.Second)
}

type RoomLimitConfig struct {
	RoomId    int   `json:"room_id"`
	WhiteList []int `json:"white_list"`
	LimitNum  int   `json:"limit_num"`
}
