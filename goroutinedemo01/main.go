package main

import (
	"fmt"
	"time"
)

func echoHello() {
	for i := 1; i <= 10; i++ {
		fmt.Println("hello go", i)
		time.Sleep(time.Second)
	}
}

func main() {
	fmt.Println("1111111")
	//获取当前系统cpu数量
	//cpuNum := runtime.NumCPU()
	//cpuNum-=1
	//runtime.GOMAXPROCS(cpuNum)

	go echoHello()

	//等待上面的echoHello执行完
	for i := 1; i <= 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello main", i)
	}
}
