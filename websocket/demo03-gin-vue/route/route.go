package route

import (
	"haimait/learn/websocket/demo03-gin-vue/controller/ws"

	"github.com/gin-gonic/gin"
)

// Init 路由初始化
func Init() {
	router := gin.Default()
	router.Use(Cors()) //解决跨域问题
	router.GET("/ping", ws.Ping)
	router.GET("/test", ws.Test)
	router.Run(":3000")
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			//c.AbortWithStatus(http.StatusNoContent)
			c.AbortWithStatus(200)
		}
		// 处理请求
		c.Next()
	}
}
