package main

import (
	"flag"
	"log"
	"net/http"
	"shuoba/hub"
	"shuoba/views"
	"time"
)











func main(){
	//orm:=model.SetEngine()
	port := flag.String("port","12345","http listen port")
	flag.Parse()

	//生成hub交换机
	hubs := hub.NewHub()
	go hubs.Run()
	//静态文件设置
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	//routing
	http.HandleFunc("/",views.Home)
	http.HandleFunc("/login",views.LoginHandle)
	http.HandleFunc("/test",views.TestHandle)
	http.HandleFunc("/servicep",views.ServicePeopleHandle)
	http.HandleFunc("/chat",views.ChatHome)
	//http.Handle("/chat",views.ValidateTokenMiddleware(http.HandlerFunc(views.ChatHome)))
	http.Handle("/admin",views.ValidateTokenMiddleware(http.HandlerFunc(views.AdminHome)))
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		hub.ServerWsSwitch(hubs,w,r)
	})
	log.Printf("start server 0.0.0.0:%v\n",*port)
	srv := &http.Server{
		Handler:nil,
		Addr:"0.0.0.0:"+*port,
		ReadTimeout:10 * time.Second,
		WriteTimeout:10 * time.Second,
	}
	err :=srv.ListenAndServe();if err !=nil{
		log.Fatal("start error",err)
	}
}
