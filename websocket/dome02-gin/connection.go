package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/gorilla/websocket"
)

type connection struct {
	ws   *websocket.Conn //ws连接器
	sendChan   chan []byte  //管道send
	data *Data //数据
}

//定义升级器，将http请求升级为ws请求
var wu = &websocket.Upgrader{
	ReadBufferSize: 1024, //在ws中指定读缓存区大小
	WriteBufferSize: 1024,  //在ws中指定写缓存区大小
	CheckOrigin: func(r *http.Request) bool { return true }, //充许跨域访问
}

//ws的回调函数
func myws(ctx *gin.Context) {
	//1.获取wd的对象
	ws, err := wu.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		return
	}
	//初始化连接对象
	c := &connection{
		ws: ws,
		sendChan: make(chan []byte, 256),
		data: &Data{},
	}
	fmt.Println("c=====>",*c)
	//在ws中注册一下
	h.register <- c
	//wd将数据读写跑起来
	go c.writer()
	c.reader()

	//回调结束以后，相当于关闭浏览器了，和正常退出一样的处理
	defer func() {
		c.data.Type = "logout"
		//用户列表删除
		userList = del(userList, c.data.User)
		c.data.UserList = userList   //删除后更新用户数据
		c.data.Content = c.data.User //删除后传播XXX已经下线了

		//数据序列化，让所有人看到XXX下线了
		dataB, _ := json.Marshal(c.data)
		h.broadcast <- dataB
		h.register <- c
	}()
}

func (c *connection) writer() {
	fmt.Println("sendChan",c.sendChan)
	for message := range c.sendChan {
		c.ws.WriteMessage(websocket.TextMessage, message) //把前端发来的消息不断的写入给发消息通知的管道里
	}
	c.ws.Close()
}

var userList = []string{}

func (c *connection) reader() {
	for {
		_, message, err := c.ws.ReadMessage() //不断的从消息通道里读数据
		if err != nil {
			h.register <- c
			break
		}
		json.Unmarshal(message, &c.data)
		fmt.Println("data:",*c.data)
		switch c.data.Type {
		case "login":
			c.data.User = c.data.Content
			c.data.From = c.data.User
			userList = append(userList, c.data.User)
			c.data.UserList = userList
			dataB, _ := json.Marshal(c.data)
			h.broadcast <- dataB
		case "user":
			c.data.Type = "user"
			dataB, _ := json.Marshal(c.data)
			h.broadcast <- dataB
		case "logout":
			c.data.Type = "logout"
			userList = del(userList, c.data.User)
			c.data.UserList = userList   //删除后更新用户数据
			c.data.Content = c.data.User //删除后传播XXX已经下线了
			dataB, _ := json.Marshal(c.data)
			h.broadcast <- dataB
			h.register <- c
		default:
			fmt.Print("========default================")
		}
	}
}

//删除用户
func del(slice []string, user string) []string {
	count := len(slice)
	if count == 0 {
		return slice
	}
	if count == 1 && slice[0] == user {
		return []string{}
	}
	//定义新的返回切片
	var mySlice = []string{}
	//删除传入切片中的 指定用户，其它用户用户放到新的切片
	for i := range slice {
		//利用索引删除用户 如果当前是切片里的最后一个元素，就删除它
		if slice[i] == user && i == count {
			return slice[:count]
		} else if slice[i] == user {
			mySlice = append(slice[:i], slice[i+1:]...)
			break
		}
	}
	fmt.Println(mySlice)
	return mySlice
}
