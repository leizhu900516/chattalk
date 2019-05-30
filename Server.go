package main

import (
	"../shuoba/hub"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

//消息结构
// msgtype 几种角色：
//			1、客服是否在线
//			2、获取历史消息
type Message struct {
	Userid string `json:"userid"`
	Destid string `json:"destid"`
	Content string `jons:"content"`
	Addtime uint64 `json:"addtime"`
	MsgType uint8 `json:"msgtype"`
}
type Hub struct {
	conn *websocket.Conn

}
//客服人员
var ServicePersonnel struct{
	Sid uint64
	online bool
}
//信息交换机
//根据不同的连接conn发送给固定的客服人员
var Switch = make(map[uint64]Hub)


var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:1024,
	WriteBufferSize:1024,
}

func chatHome( w http.ResponseWriter, r *http.Request){
	http.ServeFile(w,r,"templates/chat-page.html")
}
func adminHome( w http.ResponseWriter, r *http.Request){
	http.ServeFile(w,r,"templates/admin.html")
}


//func WsHandle(w http.ResponseWriter,r *http.Request){
//	print("新连接")
//	r.ParseForm()
//	_userid := r.Form["userid"][0]
//	userid, err := strconv.ParseUint(_userid, 10, 64)
//	_destid := r.Form["destid"][0]
//	destid, err := strconv.ParseUint(_destid, 10, 64)
//
//	//fmt.Println("userid=",userid,reflect.TypeOf(userid))
//	fmt.Println("destid=",destid,reflect.TypeOf(destid))
//	//获取socket连接
//	conn,err := upgrader.Upgrade(w,r,nil)
//	address :=conn.LocalAddr()
//	fmt.Println("ip=",address)
//
//	if _,ok :=Switch[userid];!ok{
//		Switch[userid] = Hub{conn:conn}
//	}
//	//Switch[conn]
//	fmt.Println("conn=",&conn)
//	if err !=nil{
//		log.Println(err)
//		return
//	}
//	for{
//		//接收消息
//		messageType,p,err :=conn.ReadMessage()
//		if err !=nil{
//			//连接断开
//			delete(Switch,userid)
//		}
//		var msg Message
//		if err := json.Unmarshal(p,&msg);err !=nil{
//			log.Println(err)
//		}
//		if _,ok :=Switch[msg.Userid];!ok{
//			Switch[msg.Userid] = Hub{conn:conn}
//		}
//		fmt.Println("recv=",msg.Destid)
//		if err != nil {
//			log.Println(err)
//			return
//		}
//		fmt.Println(len(Switch))
//		fmt.Println("messageType=",messageType)
//		fmt.Println("p=",p)
//
//		//发送消息给目标;对方在线
//		if v,ok := Switch[msg.Destid];ok{
//			fmt.Println(v.conn)
//			if err := v.conn.WriteMessage(messageType,p);err !=nil{
//				log.Println(err)
//				return
//			}
//		}else { //对方不在线
//			if err := conn.WriteJSON(Message{Userid:msg.Userid,Destid:msg.Destid,Content:"对方不在线aaa",Addtime:123456789});err !=nil{
//				log.Println(err)
//				return
//			}
//		}
//		//if err := conn.WriteMessage(messageType,p);err !=nil{
//		//	log.Println(err)
//		//	return
//		//}
//	}
//
//
//}
func main(){

	hubs := hub.NewHub()
	go hubs.Run()
	//静态文件设置
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	http.HandleFunc("/chat",chatHome)
	http.HandleFunc("/admin",adminHome)
	//http.HandleFunc("/ws",WsHandle)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		hub.ServerWsSwitch(hubs,w,r)
	})
	fmt.Println("start server 0.0.0.0:9090")
	err :=http.ListenAndServe("0.0.0.0:9090",nil);if err !=nil{
		log.Fatal("start error")
	}
}
