package main

import "github.com/gin-gonic/gin"

func main()  {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello golang",
		})
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/test", Test)

	r.Run("0.0.0.0:9999") // 监听并在 0.0.0.0:8080 上启动服务
}
func Test(ctx *gin.Context) {

	ctx.JSON(200, gin.H{
		"data": "test ok",
	})
}