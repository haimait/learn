package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)



type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func Test(ctx *gin.Context) {
	//读取
	var product Product
	DB.First(&product, 1) // 查询id为1的product
	fmt.Println("product", product)

	ctx.JSON(200, gin.H{
		"data": product,
	})
}

var (
	DB *gorm.DB
)

func InitMysql() (err error) {
	DB, err = gorm.Open("mysql", "root:c8e772bb8a4ef9e0@/gin01?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("open mysql failed", err)
	}

	// 全局禁用表名复数
	DB.SingularTable(true)
	// 自动迁移模式
	//DB.AutoMigrate(&Product{})
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
	// 启用Logger，显示详细日志
	DB.LogMode(true)
	DB.DB().Ping()
	return
}
func main() {
	err := InitMysql()
	if err != nil {
		panic(err)
	}
	defer DB.Close()
	// 创建
	//db.Create(&Product{Code: "L1212", Price: 1000})

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/test", Test)

	r.Run("0.0.0.0:8082") // 监听并在 0.0.0.0:8080 上启动服务

	// 读取
	//var product Product
	//db.First(&product, 1) // 查询id为1的product
	//fmt.Println("product",product)
	//
	//db.First(&product, "code = ?", "L1212") // 查询code为l1212的product
	//fmt.Println("product",product)
	//fmt.Println("product",product.CreatedAt)

	//// 更新 - 更新product的price为2000
	//db.Model(&product).Update("Price", 2000)
	//
	//// 删除 - 删除product
	//db.Delete(&product)
}
