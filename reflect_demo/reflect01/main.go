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
func Test() {
	var s Student
	s.Class = "2一1"
	Check(s)
}

func Test3() {
	var u User
	u.Name = "lish"
	u.Age = 18
	u.Gender = "男"
	Check(u)
}

//断言 取struct里的值
func Check(v interface{}) {
	//v.(User).SayName(v.(User).Name) //断言为User类型
	//判断是不是User判断
	u, ok := v.(User)
	if !ok {
		fmt.Println("reflect User failed")
		return
	}
	fmt.Println(u.Name)

}
func main() {
	//Test() //断言
	Test3() //断言
}
