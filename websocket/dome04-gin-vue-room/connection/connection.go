package connection

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/gorilla/websocket"
)

//将连接器对象初始化 ws的控制器
type hub struct {
	connection map[*connection]bool //connection 初始化连接对象 标识用户已经注册了
	broadcast chan []byte  //broadcast 从连接器(前端)发送的信息
	register chan *connection //register 从连接器注册的请求 (连接通道)
	unregister chan *connection //unregister 销毁的请求(所有的连接通道)
}
var H = hub{
	connection: make(map[*connection]bool),
	unregister: make(chan *connection),
	broadcast: make(chan []byte),
	register: make(chan *connection),
}

//每个连接通
type connection struct {
	ws   *websocket.Conn //ws连接器
	sendChan  chan []byte
	data *Data //前后端要交互发送的数据
}


type Data struct {
	Ip       string   `json:"ip"` //用户ip
	User     string   `json:"user"` //用户名
	From     string   `json:"from"` //来自哪个用户的信息
	Type     string   `json:"type"` //类型  login/user/logout
	Content  string   `json:"content"` //发送的内容
	UserList []string `json:"user_list"`  //连接上的所有用户
}
//定义升级器，将http请求升级为ws请求
var wu = &websocket.Upgrader{
	ReadBufferSize: 1024, //在ws中指定读缓存区大小
	WriteBufferSize: 1024,  //在ws中指定写缓存区大小
	CheckOrigin: func(r *http.Request) bool { return true }, //充许跨域访问
}
var ws *websocket.Conn
//ws的回调函数
func Myws(ctx *gin.Context) {
	//1.获取wd的对象
	ws, err := wu.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		return
	}
	//初始化连接对象
	c := &connection{
		ws: ws,
		sendChan : make(chan []byte, 256),
		data: &Data{},
	}
	fmt.Println("c.",c)
	//在ws中注册一下
	H.register <- c
	//wd将数据读写跑起来

	go writer(&c.sendChan, c.ws)
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
		H.broadcast <- dataB
		H.register <- c
	}()
}

//把通道的消息发给前端
func writer(sendChan *chan []byte,ws *websocket.Conn) {
	for message := range *sendChan {
		fmt.Printf("c.message222222111 %#v \n",string(message))
		ws.WriteMessage(websocket.TextMessage, message) //把前端发来的消息不断的写入给发消息通知的管道里
	}
	ws.Close()
}

var userList = []string{}

func (c *connection) reader() {
	for {
		_, message, err := c.ws.ReadMessage() //不断的从消息通道里读数据
		if err != nil {
			H.register <- c
			break
		}
		fmt.Printf("message:%#v \n",string(message))

		json.Unmarshal(message, &c.data)
		fmt.Printf("data11:%#v \n",*c.data)
		switch c.data.Type {
		case "login":
			//{\"type\":\"login\",\"content\":\"a1\"}
			c.data.User = c.data.Content
			c.data.From = c.data.User
			userList = append(userList, c.data.User)
			c.data.UserList = userList
			dataB, _ := json.Marshal(c.data)
			fmt.Printf("login--dataB:%#v \n",c.data)

			H.broadcast <- dataB
		case "user":
			c.data.Type = "user"
			dataB, _ := json.Marshal(c.data)
			fmt.Printf("user--dataB:%#v \n",c.data)
			H.broadcast <- dataB
			Test()
		case "logout":
			c.data.Type = "logout"
			userList = del(userList, c.data.User)
			c.data.UserList = userList   //删除后更新用户数据
			c.data.Content = c.data.User //删除后传播XXX已经下线了
			dataB, _ := json.Marshal(c.data)
			H.broadcast <- dataB
			H.register <- c
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

//处理ws的逻辑实现
func (h *hub) Run() {
	//监听数据管道，在后端不断处理管道数据
	for {
		//循环所有的连接，根据不同的数据管道，处理不同的逻辑
		select {
		//注册
		case c := <-h.register: //h.register
			h.connection[c] = true //标志注册了
			//组装data数据
			c.data.Ip = c.ws.RemoteAddr().String() //登录用户的ip
			c.data.Type = "handshake"              //更新类型
			c.data.UserList = userList             //用户列表
			dataB, _ := json.Marshal(c.data)       //将数据序列化一下
			c.sendChan <- dataB                    //将数据放入数据管道

		case c := <-h.unregister: //h.unregister注销用户
			if _, ok := h.connection[c]; ok { //判断map里存存在要删除的数据再删除
				delete(h.connection, c)
				close(c.sendChan)
			}
		case data := <-h.broadcast:  //h.broadcast 处理前端传来的消息通道
			for c := range h.connection { //h.connection循环当前所有在线用户的连接
				select {
				case c.sendChan <- data: //将消息同步给所有的数据管道，通知所有的人
				default:
					//防止死循环
					delete(h.connection, c) //h.connecttions 如果一直没有接收数据，就删除用户的连接，以免在这里占用资源
					close(c.sendChan) //c.send判断管道
				}
			}
		}
	}
}


func Test(){
	//初始化连接对象
	var c = &connection{}
	c.data.Type = "user"
	c.data.Content = "test"
	dataB, _ := json.Marshal(c.data)       //将数据序列化一下
	H.broadcast <- dataB
	//c.sendChan <- dataB
	//writer(&c.sendChan , ws)
}