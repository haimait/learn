package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"haimait/learn/pachong/guoshu_goodsApi/app/model/mall_goods"
)

func main() {
	s := g.Server()

	// 常规注册
	group := s.Group("/")

	// 登录页面3
	group.GET("/ping", func(r *ghttp.Request) {
		glog.Println("ping")
		//data := g.Map{
		//	"title": "登录页面",
		//}
		//g.Log().Info("[default]info")
		//g.Log().Info(data)
		//g.Log().Debug(data)
		//g.Log().Warning("[default]Warning")
		//g.Log().Error("[default]Error")

		r.Response.WriteJson(g.Map{
			"code": 0,
			"msg":  "pong",
		})
		r.Exit()

		//r.Response.WriteJson(g.Map{
		//	"code": -1,
		//	"msg":  e.Error(),
		//})
		//r.Exit()
	})



	// 用户对象
	//type User struct {
	//	Username string `gvalid:"username     @required|length:5,16#请输入用户名称|用户名称长度非法,5-16个字符"`
	//	Password string `gvalid:"password     @required|length:31,33#请输入密码|密码长度非法,31-33个字符"`
	//}

	//获取分类下的商品
	group.GET("/getGoodsList", mall_goods.GetGoodsList)



	//商品详情
	group.GET("/getGoodsDetail", mall_goods.GetGoodsDetail)
	//group.GET("/getGoodsDetail", func(r *ghttp.Request) {
	//	glog.Println("getGoodsDetail")
	//	var path = "https://api.jiafuminkang.com/app/Product/bulkDetailXiao?goods_id=14146&mer_id=&sessionId="
	//
	//	// GET请求
	//	if response, err := ghttp.Get(path); err != nil {
	//		panic(err)
	//	} else {
	//		defer response.Close()
	//		g.Log().Line().Info(response.ReadAllString())
	//	}
	//
	//
	//
	//
	//	//if response, err := ghttp.Post(path); err != nil {
	//	//	panic(err)
	//	//} else {
	//	//	defer response.Close()
	//	//	g.Log().Line().Info(response.ReadAllString())
	//	//}
	//
	//	//data := g.Map{
	//	//	"title": "登录页面",
	//	//}
	//	//g.Log().Info("[default]info")
	//	//g.Log().Info(data)
	//	//g.Log().Debug(data)
	//	//g.Log().Warning("[default]Warning")
	//	//g.Log().Error("[default]Error")
	//
	//	r.Response.WriteJson(g.Map{
	//		"code": 0,
	//		"msg":  "ok",
	//	})
	//	r.Exit()
	//
	//	//r.Response.WriteJson(g.Map{
	//	//	"code": -1,
	//	//	"msg":  e.Error(),
	//	//})
	//	//r.Exit()
	//})



	//s.SetHTTPSPort(8080)
	s.Run()
}
