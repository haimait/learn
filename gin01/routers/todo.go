package routers

import (
	"github.com/gin-gonic/gin"
	"haimait/learn/gin01/controller"
)

func InitTodoRouter(Router *gin.RouterGroup) {
	TodoRouter := Router.Group("")
	{
		// 添加
		TodoRouter.POST("/todo", controller.CreateTodo)
		// 查看所有的待办事项
		TodoRouter.GET("/todo", controller.GetTodoList)
		// 修改某一个待办事项
		TodoRouter.PUT("/todo/:id", controller.UpdateATodo)
		// 删除某一个待办事项
		TodoRouter.DELETE("/todo/:id", controller.DeleteATodo)
	}
}
