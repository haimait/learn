package main

import (
	"fmt"
	"math/rand"
	"time"
)

//func a() bool {
//	return rand.Intn(10)>5
//}
//
//func test1(){
//	expr := 1==1
//	switch expr {
//	case a():
//		fmt.Println(111)
//	case a():
//		fmt.Println(222)
//	case a():
//		fmt.Println(333)
//	}
//}
func main() {
	t:=rand1()
	switch rand1() {
	case t:
		fmt.Println(11111)
	case t:
		fmt.Println(222222)
	}
}

func rand1()bool{
	//将时间戳设置成种子数
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(10)>5
}