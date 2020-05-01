package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"haimait/learn/gin01/dao"
	"haimait/learn/gin01/models"
	"haimait/learn/gin01/routers"
)

// Inject fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt` into models `User`
// 将 `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`字段注入到`User`模型中
type User struct {
	gorm.Model
	Name string
}


func Pong(c *gin.Context)  {
	c.JSON(200,gin.H{
		"message":"pong",
	})
}

func main() {

	// 创建数据库
	// sql: CREATE DATABASE bubble;
	// 连接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer dao.Close()  // 程序退出关闭数据库连接


	r := routers.SetupRouter()

	//r:=gin.Default()
	//r.GET("/ping", Pong)
	r.Run("0.0.0.0:8080")
}
