package main

import "encoding/json"

//将连接器对象初始化 ws的控制器
var h = hub{
	connection: make(map[*connection]bool),
	unregister: make(chan *connection),
	broadcast: make(chan []byte),
	register: make(chan *connection),
}

type hub struct {
	connection map[*connection]bool //connection 初始化连接对象 标识用户已经注册了
	broadcast chan []byte  //broadcast 从连接器(前端)发送的信息
	register chan *connection //register 从连接器注册的请求
	unregister chan *connection //unregister 销毁的请求
}

//处理ws的逻辑实现
func (h *hub) run() {
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