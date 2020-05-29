package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter() //创建一个路由
	go h.run()                //ws控制器 监听h里的管道，不断的处理管道里的数据（用户的注册，注销，发送消息等），进行同步数据

	//指定ws回调函数，监听前端的访问路由
	router.HandleFunc("/ws", myws)
	//开启服务端监听
	if err := http.ListenAndServe("127.0.0.1:8080", router); err != nil {
		fmt.Println("err:", err)
	}
}
