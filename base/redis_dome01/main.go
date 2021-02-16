package main

import (
	"fmt"
	"haimait/learn/base/redis_dome01/utils"
)



func main(){
	test1()
}

func test1()  {
	utils.RedisSet("age",18,0)
	get, err := utils.RedisGet("age2")
	if err !=nil {
		fmt.Println("err:",err)
		return
	}
	fmt.Println(get)
}