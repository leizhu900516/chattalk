package hub

import (
	"fmt"
	"log"
	"time"
	"encoding/json"
)

type Hub struct {
	client map[*Client] bool
	register chan *Client
	unregister chan *Client
	hubport map[string] *Client //转发到指定端口=用户id
	broadcast chan string //客户退出的广播通道
}





func NewHub() *Hub{
	return &Hub{
		client: make(map[*Client]bool),
		register: make(chan *Client),
		unregister: make(chan *Client),
		hubport: make(map[string]*Client),
		broadcast: make(chan string,2),
	}
}


//注册新的client
//删除掉旧的client
func (hub *Hub) Run(){
	for {
		select {
		case client := <- hub.register:
			fmt.Println("register",client)
			hub.client[client] = true

		case client := <- hub.unregister:
			fmt.Println("有客户端关闭",client)
			hub.broadcast <- client.uid
			if _,ok := hub.client[client];ok{
				//todo 客户对应哪个客服
				// 像客服的队列发送一个客户退出的消息  msgClientClose
				// 客服前台根据标示，标记目标客服下线。？时间段内删除左侧下线的客户列表(避免客户过多)
				log.Println("删除客户",client.uid)
				delete(hub.client,client)
				close(client.send)
				delete(hub.hubport,client.uid)

			}else {
				fmt.Println("客服端不在hub中")
			}
			fmt.Println("清除成功")
		case uid := <- hub.broadcast:
			//通告所有的客服

			kefuid := "1000"
			if uid != kefuid{
				fmt.Println("通知客服用户离线",uid)
				responsedata,_ := json.Marshal(Message{Userid:uid,Destid:kefuid,Content:"用户离线",Addtime:time.Now().String()[:19],MsgType:msgClientClose})
				hub.hubport[kefuid].send <- []byte(responsedata)
			}

			//for key,value := range *ServicePersionMap{
			//	fmt.Println(key)
			//	fmt.Println(value)
			//}
		}
	}
}





