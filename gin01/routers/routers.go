package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine{
	Router := gin.Default()
	// 告诉gin框架模板文件引用的静态文件去哪里找
	Router.Static("/public", "public")
	// 告诉gin框架去哪里找模板文件
	//Router.LoadHTMLGlob("templates/*")
	//Router.GET("/", controller.IndexHandler)
	//Router.GET("test", controller.TestHandler)


	Router.MaxMultipartMemory = 8 << 20 // 8 MiB 提交文件时默认的内存限制是32 MiB

	// 方便统一添加路由组前缀 多服务器上线使用
	ApiGroup := Router.Group("")

	//注册路由
	InitTodoRouter(ApiGroup)

	fmt.Println("router register success")
	return Router
}
