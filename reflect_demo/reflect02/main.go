package main

import "fmt"

type User struct {
	Name   string
	Age    int
	Gender string
}
type Student struct {
	Class string
	User
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

func TypeDemo02()  {
	var s Student
	s.Class = "2一1"
	Check2(s)
}
// type 获取类型 判断传入的类型是否是我们想要的类型
func Check2(v interface{}) {
	switch v.(type){
	case User:
		fmt.Println(v.(User).Name) //lish
	case Student:
		fmt.Println(v.(Student).Class) //2一1
	}
}
func main() {
	TypeDemo01()//type获取类型
	TypeDemo02()//type获取类型
}
