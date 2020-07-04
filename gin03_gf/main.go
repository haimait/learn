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
	r.GET("valid", Valid) //验证器demo
	r.GET("gfconfig", GfConfig) //配置demo
	r.GET("gflog", GfLog) //日志demo
	r.GET("/json", func(c *gin.Context) {
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

//配置demo
func GfConfig(ctx *gin.Context) {
	name:=g.Config().GetString("name")
	fmt.Println(name)  //hello world!
	g.Dump(name)  //"hello world!"

	c := g.Cfg()
	// 分组方式
	fmt.Println(c.Get("redis.cache")) //127.0.0.1:6379,1
	g.Dump(c.Get("redis.cache")) //"127.0.0.1:6379,1"

	// 数组方式：test2
	port:=c.GetInt("database.default.0.port")  //type:int  value:3306
	fmt.Printf("type:%T  value:%#v \n",port,port) //取2号单元下name
	g.Dump(port) //3306
}

//日志demo
func GfLog(ctx *gin.Context) {

	g.Log().Debug("[default]Debug")
	g.Log().Info("[default]info")
	g.Log().Line().Info("[你来了]info")
	g.Log().Warning("[default]Warning")
	g.Log().Error("[default]Error")
	// 异常
	g.Log().Panic("this is panic！")
}

type test01DataInput struct {
	//Name     string `json:"name" form:"name" v:"required|length:6,16#账号不能为空|账号长度应当在:min到:max之间"`
	//Name     string `json:"name" form:"name" v:"required|length:6,16"`
	Name     string `json:"name" form:"name" valid:"name  @required|length:6,30#name不能为空|name长度应当在:6到:30之间"`
	PassWord string `json:"password" form:"password" v:"required|length:6,16#密码不能为空|账号长度应当在:6到:16之间"`
}

//验证器
func Valid(ctx *gin.Context) {
	var t test01DataInput
	_ = ctx.ShouldBindQuery(&t)
	// 输入参数检查
	if e := gvalid.CheckStruct(t, nil); e != nil {
		fmt.Println("1111111", errors.New(e.FirstString()))
		g.Dump(e.Maps())
		//ctx.JSON(http.StatusOK, e.Maps())
		ctx.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "error",
			//"data":e.Maps(),
			//"data":e.String(),
			"data": e.FirstString(),
		})
		return
	}

	ctx.JSON(http.StatusOK, t)
}
