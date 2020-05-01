package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name   string
	Age    int
	Gender string
}
type Student struct {
	Class string
	User  User
}

func (u User) SayName(name string) {
	fmt.Println("my name is ", name)
}

func TypeDemo01() {
	var u User
	u.Name = "lish"
	u.Age = 18
	u.Gender = "男"
	Check(u)

}

// 反射 调用结构体里的方法
func Check(data interface{}) {
	v:= reflect.ValueOf(data)
	m := v.Method(0)
	m.Call([]reflect.Value{reflect.ValueOf("大奇")})
}


func main() {
	TypeDemo01() //type获取类型
}
