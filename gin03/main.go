package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gvalid"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("test01", test01)
	r.GET("/JSON", func(c *gin.Context) {
		//data := map[string]interface{}{
		//	"foo": "bar",
		//}
		data1 := gin.H{
			"foo": "bar",
		}
		c.JSON(http.StatusOK, data1)
	})

	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8080")
}

type test01DataInput struct {
	//Name     string `json:"name" form:"name" v:"required|length:6,16#账号不能为空|账号长度应当在:min到:max之间"`
	//Name     string `json:"name" form:"name" v:"required|length:6,16"`
	Name  string `json:"name" form:"name" valid:"name  @required|length:6,30"`
	PassWord string `json:"password" form:"password" v:"required|length:6,16#密码不能为空|账号长度应当在:min到:max之间"`
}

func test01(ctx *gin.Context) {
	var t test01DataInput
	_ = ctx.ShouldBindQuery(&t)
	// 输入参数检查
	if e := gvalid.CheckStruct(t, nil); e != nil {
		fmt.Println("1111111",errors.New(e.FirstString()))
		g.Dump(e.Maps())
		ctx.JSON(http.StatusOK, e.Maps())
		return
	}

	ctx.JSON(http.StatusOK, t)
}
