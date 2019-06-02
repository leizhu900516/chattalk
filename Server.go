package main

import (
	"shuoba/hub"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"shuoba/views"
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




func main(){

	hubs := hub.NewHub()
	go hubs.Run()
	//静态文件设置
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	http.HandleFunc("/",views.Home)
	http.HandleFunc("/login",views.LoginHandle)
	//http.HandleFunc("/chat",views.ChatHome)
	http.Handle("/chat",views.ValidateTokenMiddleware(http.HandlerFunc(views.ChatHome)))
	http.Handle("/admin",views.ValidateTokenMiddleware(http.HandlerFunc(views.AdminHome)))
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		hub.ServerWsSwitch(hubs,w,r)
	})
	fmt.Println("start server 0.0.0.0:12345")
	err :=http.ListenAndServe("0.0.0.0:12345",nil);if err !=nil{
		log.Fatal("start error")
	}
}
