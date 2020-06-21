package ws

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Data struct {
	ToUser  string `json:"touser"`
	Type    string `json:"type"`
	Content string `json:"content"`
}

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// type data struct {
// 	Msg string `json:"msg"`
// }
var sendChan chan []byte     //管道send

type connection struct {
	ws       *websocket.Conn //ws连接器
	data     *Data           //数据
}

//定义升级器，将http请求升级为ws请求
var wu = &websocket.Upgrader{
	ReadBufferSize:  1024,                                       //在ws中指定读缓存区大小
	WriteBufferSize: 1024,                                       //在ws中指定写缓存区大小
	CheckOrigin:     func(r *http.Request) bool { return true }, //充许跨域访问
}

// Ping webSocket请求Ping 返回pong
func Ping(ctx *gin.Context) {
	//升级get请求为webSocket协议
	ws, err := wu.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		return
	}
	//初始化连接对象
	c := &connection{
		ws:       ws,
		data:     &Data{},
	}
   sendChan = make(chan []byte, 256)

	go c.writer(&sendChan)
	c.reader(&sendChan)
	//defer ws.Close()

}

func (c *connection) reader(*chan []byte) {
	for {
		// 读取ws中的数据
		mt, message, err := c.ws.ReadMessage()
		if err != nil {
			// 客户端关闭连接时也会进入
			fmt.Println(err)
			break
		}
		// msg := &data{}
		// json.Unmarshal(message, msg)
		// fmt.Println(msg)
		fmt.Println(mt)
		fmt.Println(message)
		fmt.Println(string(message))
		// 如果客户端发送ping就返回pong,否则数据原封不动返还给客户端
		if string(message) == "ping" {
			data,_:=json.Marshal(Data{"ping1","ping1","ping1"})
			message = []byte(data)
			//c.writer(message)
		}else{
			data,_:=json.Marshal(Data{"wangwu","send",string(message)})
			message = []byte(data)
		}
		sendChan <- message
	}

}

//发送消息
//func sendMessage(messageType int, message []byte)  {
//	ws.WriteMessage(messageType, message)
//}
func (c *connection) writer(sendChan *chan []byte) {
	fmt.Println("sendChan",sendChan)
	for message := range *sendChan {
		c.ws.WriteMessage(websocket.TextMessage, message) //把前端发来的消息不断的写入给发消息通知的管道里
	}
	c.ws.Close()
}

func Test(ctx *gin.Context) {
	data,_:=json.Marshal(Data{"lisi","massage","test"})
	message := []byte(data)
	sendChan <- message
	ctx.JSON(200,gin.H{"code":200,"msg":"ok"})
}
