package main

import (
	"github.com/gin-gonic/gin"
	"haimait/learn/websocket/dome02-gin/connection"
)

func main() {

	go connection.H.Run()  //ws控制器 监听h里的管道，不断的处理管道里的数据（用户的注册，注销，发送消息等），进行同步数据

	r := gin.Default()
	r.GET("/ws",connection.Myws)
	r.Run("127.0.0.1:8080")
}
