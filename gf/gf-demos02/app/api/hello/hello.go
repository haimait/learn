package hello

import (
    "fmt"
    "gf01/app/model/users"
    "github.com/gogf/gf/net/ghttp"
)

// Hello is a demonstration route handler for output "Hello World!".
func Hello(r *ghttp.Request) {
    r.Response.Writeln("Hello World!")
    userInfo, err := users.FindOne("username = ?", "admin")
    if err !=nil {
        //glog.Error(err)
        fmt.Println(err)
        r.Response.Writefln("err")
        r.Exit()
    }
    r.Response.Writefln(userInfo.NickName)
}
