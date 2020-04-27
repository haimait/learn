package main

import(
	"fmt"
	"runtime"
	"time"
)

func echoHello()  {
	for i:=1;i<=10 ;i++  {
		fmt.Println("hello go",i)
		time.Sleep(time.Second)
	}
}

func main(){
	fmt.Println("1111111")
	cpuNum:=runtime.NumCPU()
	fmt.Println(cpuNum)

	go echoHello()
	for i:=1;i<=10 ;i++  {
		time.Sleep(time.Second)
		fmt.Println("hello main",i)
	}
}