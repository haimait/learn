package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/glog"
)

func main() {
	s := g.Server()
	// 测试日志
	s.BindHandler("/welcome", func(r *ghttp.Request) {
		glog.Info("你来了！")
		glog.Error("你异常啦！")
		r.Response.Write("哈喽世界！")
	})
	// 测试日志
	s.BindHandler("/welcome2", func(r *ghttp.Request) {
		path := "logs"
		glog.SetPath(path)
		glog.Println("日志内容")
		glog.Info("你来了！")
		glog.Error("你异常啦！")
		list, err := gfile.ScanDir(path, "*") //看logs目录下的所有文件
		g.Dump(err)
		g.Dump(list)
		//打印调用行号
		glog.Line().Println("this is the short file name with its line number") //相对路径
		glog.Line(true).Println("lone file name with line number") //绝对路径
		r.Response.Write("哈喽世界！")
	})
	// 测试日志
	s.BindHandler("/welcome3", func(r *ghttp.Request) {
		// 对应默认配置项 logger，默认default
		g.Log().Debug("[default]Debug")
		g.Log().Info("[default]info")
		g.Log().Line().Info("[你来了]info")
		g.Log().Warning("[default]Warning")
		g.Log().Error("[default]Error")
	})
	// 异常处理
	s.BindHandler("/panic", func(r *ghttp.Request) {
		glog.Line().Println("println_123")
		glog.Panic("123")
		//r.Response.Writeln("This is panic!")
		r.Response.WriteJson(g.Map{
			"code": -1,
			"msg":  "This is panic!",
		})
		r.Exit()
	})
	// post请求
	s.BindHandler("POST:/hello", func(r *ghttp.Request) {
		r.Response.Writeln("Hello World!")
	})
	s.Run()
}
