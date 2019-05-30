package hub

import (
	"fmt"
	"log"
)

type Hub struct {
	client map[*Client] bool
	register chan *Client
	unregister chan *Client
	port map[string] *Client //转发到指定端口
}


func NewHub() *Hub{
	return &Hub{
		client: make(map[*Client]bool),
		register: make(chan *Client),
		unregister: make(chan *Client),
		port: make(map[string]*Client),

	}
}


//注册新的client
//删除掉旧的client
func (hub *Hub) Run(){
	for {
		select {
		case client := <- hub.register:
			hub.client[client] = true
		case client := <- hub.unregister:
			fmt.Println("有客户端关闭")
			if _,ok := hub.client[client];ok{
				log.Println("删除客户",client.uid)
				delete(hub.client,client)
				close(client.send)
				delete(hub.port,client.uid)

			}

		}
	}
}





