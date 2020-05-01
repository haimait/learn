package main

import (
	"fmt"
	"time"
)

var (
	myMap = make(map[int]int, 10)
)

func Test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= n
	}
	myMap[n] = res
}

func main() {

	for i := 1; i <= 10; i++ {
		go Test(i)
	}
	time.Sleep(time.Second * 10)

	for i, v := range myMap {
		fmt.Printf("Map[%d]=%d\n", i, v)
	}
	//输出结果
	/*
		Map[7]=823543
		Map[9]=387420489
		Map[10]=10000000000
		Map[1]=1
		Map[2]=4
		Map[6]=46656
		Map[8]=16777216
		Map[3]=27
		Map[4]=256
		Map[5]=3125
	*/

}
