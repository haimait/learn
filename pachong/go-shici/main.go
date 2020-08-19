package main

import (
	"fmt"
	"github.com/markusleevip/go-shici/gofish"
	"github.com/markusleevip/go-shici/handle"

)

func main() {
	authors := "https://so.gushiwen.org/authors/"

	h := handle.AuthorHandle{}
	fish:=gofish.NewGoFish()
	request,err := gofish.NewRequest("GET",authors,gofish.UserAgent,&h,nil)
	//添加header头信息
	request.Headers.Add("token", "123")

	if err!=nil{
		fmt.Println(err)
		return
	}
	fish.Request = request
	fish.Visit()


}
