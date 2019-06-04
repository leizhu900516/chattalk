package main

import (
	"fmt"
	"log"
	"net/http"
	"shuoba/hub"
	"shuoba/views"
)











func main(){

	hubs := hub.NewHub()
	go hubs.Run()
	//静态文件设置
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	http.HandleFunc("/",views.Home)
	http.HandleFunc("/login",views.LoginHandle)
	http.HandleFunc("/chat",views.ChatHome)
	//http.Handle("/chat",views.ValidateTokenMiddleware(http.HandlerFunc(views.ChatHome)))
	http.Handle("/admin",views.ValidateTokenMiddleware(http.HandlerFunc(views.AdminHome)))
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		hub.ServerWsSwitch(hubs,w,r)
	})
	fmt.Println("start server 0.0.0.0:12345")
	err :=http.ListenAndServe("0.0.0.0:12345",nil);if err !=nil{
		log.Fatal("start error")
	}
}
