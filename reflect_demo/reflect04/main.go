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
	Check2(u)

}

func TypeDemo02() {
	u := User{"lish", 19, "男"}

	var s Student
	s.Class = "2一1"
	s.User = u
	Check(s)
	Check2(&s)
}

// Kind()获取类型
func Check(data interface{}) {
	t := reflect.TypeOf(data)
	k := t.Kind()
	fmt.Println(k) //struct
	if k == reflect.Struct {
		fmt.Println("I am struct") //I am struct
	}
}

// 通过elem（）反射修改结构体里的值
func Check2(data interface{}) {
	v:= reflect.ValueOf(data)

	e := v.Elem()
	fmt.Println(e)//{2一1 {lish 19 男}}
	fmt.Printf("type:%T value:%#v \n",e.FieldByName("User"),e.FieldByName("User")) //{lish 19 男}

	//反射修改结构体里的值
	e.FieldByName("Class").SetString("3.2 class")
	fmt.Println(data) //&{3.2 class {lish 19 男}}
}



func main() {
	TypeDemo02() //type获取类型
}
