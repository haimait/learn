package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"haimait/learn/gin01/initialize"
	"haimait/learn/gin01/routers"
)

// Inject fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt` into models `User`
// 将 `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`字段注入到`User`模型中
type User struct {
	gorm.Model
	Name string
}

func Pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

type System struct {
	Addr string `json:"addr"`
	Port string `json:"port"`
}

func NewSystem() *System {
	return &System{
		Addr: "127.0.0.1",
		Port: "8866",
	}
}

func main() {

	// 创建数据库
	initialize.InitMySQL()

	//创建router
	r := routers.InitRouter()

	newSystem := NewSystem()

	fmt.Printf(`欢迎使用 Gin 
默认后端文件运行地址:http://127.0.0.1:%s
	`, newSystem.Port)

	r.Run(newSystem.Addr + ":" +newSystem.Port)
}
