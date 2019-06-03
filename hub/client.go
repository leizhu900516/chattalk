package hub

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"reflect"
	"time"
)

type Message struct {
	Userid string `json:"userid"`
	Destid string `json:"destid"`
	Content string `jons:"content"`
	Addtime string `json:"addtime"`
	MsgType uint8 `json:"msgtype"`
}

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)
type Client struct {
	hub *Hub
	conn *websocket.Conn
	send chan []byte
	uid string
}


func (c *Client) readPump(hub *Hub){
	fmt.Println("groutine读",hub)
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongWait));
		return nil
	})
	for {
		_,message,err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		var  msg Message
		if err := json.Unmarshal(message,&msg);err !=nil{
			log.Println(err)
		}
		//发送给目标用户的队列里
		//fmt.Println(c.conn)
		fmt.Printf("%v\n",msg)

		//c.send <- message
		fmt.Println(">>>>>",msg.Destid)
		fmt.Println(">>>>>",hub)
		fmt.Println(">>>>>",reflect.TypeOf(msg.Destid))
		if cl,ok := hub.port[msg.Destid];ok{
			fmt.Printf("%v在线",msg.Destid)
			cl.send <- message
		}else {
			responsedata,_ := json.Marshal(Message{Userid:msg.Destid,Destid:msg.Userid,Content:"不在线",Addtime:time.Now().String()[:19]})
			hub.port[msg.Userid].send <- []byte(responsedata)
			fmt.Println("不在线")
		}


	}

}

func (c *Client) writePump(hub *Hub){
	fmt.Println("groutine写")
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case message,ok := <- c.send:
			fmt.Println("send\n")
			//fmt.Println(c.conn)
			fmt.Printf("%v\n",message)
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok{
				c.conn.WriteMessage(websocket.CloseMessage,[]byte{})
				return
			}
			w ,err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil{
				return
			}
			w.Write(message)
			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage,nil);err!=nil{
				return
			}
		}
	}
}


func ServerWsSwitch(hub *Hub,w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	_userid := r.Form["userid"][0]
	//_destid := r.Form["destid"][0]
	fmt.Println(_userid)
	fmt.Println(reflect.TypeOf(_userid))

	conn,err := upgrader.Upgrade(w,r,nil)
	if err !=nil{
		log.Println(err)
		return
	}
	client := &Client{conn:conn,hub:hub,send:make(chan []byte,1000),uid:_userid}
	client.hub.register<-client
	hub.port[_userid] = client
	fmt.Println(_userid+"已经注册")
	go client.writePump(hub)
	go client.readPump(hub)
}