package main

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gsession"
	"github.com/gogf/gf/util/gvalid"
	"golang.org/x/crypto/bcrypt"
)

const SessionUser = "SessionUser"

// 认证中间件
func MiddlewareAuth(r *ghttp.Request) {
	if r.Session.Contains(SessionUser) {
		r.Middleware.Next()
	} else {
		// 获取用错误码
		r.Response.WriteJson(g.Map{
			"code": 403,
			"msg":  "您访问超时或已登出",
		})
	}
}
func main() {
	s := g.Server()

	// 设置存储方式
	sessionStorage := g.Config().GetString("SessionStorage")
	if sessionStorage == "redis" {
		s.SetSessionStorage(gsession.NewStorageRedis(g.Redis()))
		s.SetSessionIdName(g.Config().GetString("server.SessionIdName"))
	} else if sessionStorage == "memory" {
		s.SetSessionStorage(gsession.NewStorageMemory())
	}
	// 常规注册
	group := s.Group("/")
	//// 登录页面
	//group.GET("/", func(r *ghttp.Request) {
	//	r.Response.WriteTpl("index.html", g.Map{
	//		"title": "登录页面",
	//	})
	//})
	//
	//// 登录页面2
	//group.GET("/login2", func(r *ghttp.Request) {
	//	r.Response.WriteTpl("index2.html", g.Map{
	//		"title": "登录页面",
	//	})
	//})

	// 登录页面3
	group.GET("/login3", func(r *ghttp.Request) {
		glog.Println("login3日志内容")
		data:=g.Map{
			"title": "登录页面",
		}
		g.Log().Info("[default]info")
		g.Log().Info(data)
		g.Log().Debug(data)
		g.Log().Warning("[default]Warning")
		g.Log().Error("[default]Error")
		r.Response.WriteTpl("index3.html", g.Map{
			"title": "登录页面",
		})
	})

	// 用户对象
	type User struct {
		Username string `gvalid:"username     @required|length:5,16#请输入用户名称|用户名称长度非法,5-16个字符"`
		Password string `gvalid:"password     @required|length:31,33#请输入密码|密码长度非法,31-33个字符"`
	}

	group.POST("/login", func(r *ghttp.Request) {

		username := r.GetString("username")
		password := r.GetString("password")
		fmt.Println("username",username)
		fmt.Println("password",password)
		// 使用结构体定义的校验规则和错误提示进行校验
		if e := gvalid.CheckStruct(User{username, password}, nil); e != nil {
			r.Response.WriteJson(g.Map{
				"code": -1,
				"msg":  e.Error(),
			})
			r.Exit()
		}
		fmt.Println("333333333")
		record, err := g.DB().Table("sys_user").Where("login_name = ? ", username).One()
		fmt.Println("record",record)
		// 查询数据库异常
		if err != nil {
			glog.Error("查询数据错误", err)
			r.Response.WriteJson(g.Map{
				"code": -1,
				"msg":  "查询失败",
			})
			r.Exit()
		}
		// 帐号信息错误
		if record == nil {
			r.Response.WriteJson(g.Map{
				"code": -1,
				"msg":  "帐号信息错误",
			})
			r.Exit()
		}

		// 直接存入前端传输的
		successPwd := record["password"].String()
		comparePwd := password

		// 加盐密码
		// salt := "123456"
		// comparePwd, _ = gmd5.EncryptString(comparePwd + salt)

		// bcrypt验证
		err = bcrypt.CompareHashAndPassword([]byte(successPwd), []byte(comparePwd))

		//if comparePwd == successPwd {
		if err == nil {
			// 添加session
			r.Session.Set(SessionUser, g.Map{
				"username": username,
				"realName": record["real_name"].String(),
			})
			r.Response.WriteJson(g.Map{
				"code": 0,
				"msg":  "登录成功",
			})
			r.Exit()
		}

		r.Response.WriteJson(g.Map{
			"code": -1,
			"msg":  "登录失败",
		})
	})

	// 用户组
	userGroup := s.Group("/user")
	userGroup.Middleware(MiddlewareAuth)
	// 列表页面静太页面
	userGroup.GET("/index3", func(r *ghttp.Request) {
		fmt.Println("index3")
		r.Response.WriteTpl("user_index3.html")

	})


	//列表返回json数据
	userGroup.GET("/list", func(r *ghttp.Request) {
		path := "logs"
		glog.SetPath(path)
		glog.Println("日志内容")
		list, err := gfile.ScanDir(path, "*")
		g.Dump(err)
		g.Dump(list)
		//glog.Info("你来了！")
		//glog.Error("你异常啦！")
		respData := g.Map{
			"code":  0,
			"msg":   "ok",
			"title": "用户信息列表页面",
			"dataList": g.List{
				g.Map{
					"date":    "2020-04-01",
					"name":    "朱元璋",
					"address": "江苏110号",
				},
				g.Map{
					"date":    "2020-04-02",
					"name":    "徐达",
					"address": "江苏111号",
				},
				g.Map{
					"date":    "2020-04-03",
					"name":    "李善长",
					"address": "江苏112号",
				},
			},
		}
		r.Response.WriteJson(respData)

	})

	// 登出接口
	group.POST("/logout", func(r *ghttp.Request) {
		// 删除session
		r.Session.Remove(SessionUser)
		r.Response.WriteJson(g.Map{
			"code": 0,
			"msg":  "登出成功",
		})
	})

	//生成秘钥文件
	//openssl genrsa -out server.key 2048
	//生成证书文件
	//openssl req -new -x509 -key server.key -out server.crt -days 365
	s.EnableHTTPS("config/server.crt", "config/server.key")
	s.SetHTTPSPort(8080)
	s.SetPort(8199)
	s.Run()
}
