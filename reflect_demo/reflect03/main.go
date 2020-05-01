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
	User User
}

func (u User) SayName(name string) {
	fmt.Println("my name is ", name)
}

func TypeDemo02()  {
	var s Student
	s.Class = "2一1"
	s.User = User{"lish",19,"男"}
	//Check2(s)
	Check3(s)
	//Check4(s)
}
// reflect.ValueOf() ，获取输入参数接口中的数据的值
// reflect.TypeOf()  ，动态获取输入参数接口中的值的类型
func Check2(v interface{}) {
	t := reflect.TypeOf(v)
	value := reflect.ValueOf(v)
	fmt.Println(t) // main.Student
	fmt.Println(value) //{2一1 {lish 19 男}}
	//fmt.Println(t.NumField()) //2 返回结构体里的单元个数
}

//循环结构体里的值
func Check3(v interface{}) {
	t := reflect.TypeOf(v)
	value := reflect.ValueOf(v)
	fmt.Println(t) // main.Student
	fmt.Println(value) //{2一1 {lish 19 男}}
	//fmt.Println(t.NumField()) //2 返回结构体里的单元个数

	//循环结构体里的值
	for i:=0;i<t.NumField();i++{
		fmt.Println(value.Field(i)) //打印当前循环结构体里的值 2一1  {lish 19 男}
		//fmt.Println(t.Field(i)) //打印当前循环结构体里的值 2一1  {lish 19 男}
	}
}

//按index层级取struct里的值（匿名字段的时候可以这样取）
//按name名称拿stuct里的值（有字段名时可以用name名字取对应的值）
func Check4(v interface{}) {
	t := reflect.TypeOf(v)
	value := reflect.ValueOf(v)
	fmt.Println(t) // main.Student
	fmt.Println(value) //{2一1 {lish 19 男}}
	//fmt.Println(t.NumField()) //2 返回结构体里的单元个数

	//按name名称拿stuct里的值（有字段名时可以用name名字取对应的值）
	fmt.Println(value.FieldByName("Class")) //2一1
	fmt.Println(value.FieldByName("User")) //{lish 19 男}

	//按index层级取struct里的值（匿名字段的时候可以这样取）
	fmt.Println(value.Field(1)) // 取1号单元 {lish 19 男}
	fmt.Println(value.FieldByIndex([]int{1,0})) // 取1号单元下的0号单元 lish
}

func main() {
	TypeDemo02()//type获取类型
}
