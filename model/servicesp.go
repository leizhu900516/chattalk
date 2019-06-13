package model

import (
	"fmt"
	"log"
	"time"
)

//客服模型
//
type Servicesp struct {
	Id          int 		`form:"id" xorm:"int(11) pk notnull autoincr"`
	Account 	string 		`form:"account" xorm:"varchar(50) not null"`
	Sid 		string		`form:"sid" xorm:"varchar(20) not null"`
	Addtime 	int64 		`form:"addtime" xorm:"int(11) not null"`
	Password 	string 		`form:"password" xorm:"varchar(20) not null"`
	Online 		int			`form:"online" xorm:"int(1) not null default 1"`
	Realname 	string 		`form:"realname" xorm:"varchar(50) "`
}
//type Servicesp struct {
//	Id          int 		`xorm:"int(11) pk notnull autoincr"`
//	Account 	string 		`xorm:"varchar(50) not null"`
//	Sid 		string		`xorm:"varchar(20) not null"`
//	Addtime 	int64 		`xorm:"int(11) not null"`
//	Online 		int			`xorm:"int(1) not null default 1"`
//	Password 	string 		`xorm:"varchar(20) not null"`
//	Realname 	string 		`xorm:"varchar(50) "`
//}


//获取在线的客服列表提供给轮询算法
func (s *Servicesp) GetOnlineS() []Servicesp{
	var services []Servicesp
	fmt.Println("bbbb")
	fmt.Println("aaaaa",Orm.Ping())
	has ,err := Orm.Table("servicesp").Where("online = ?",1).Get(&services)
	if err !=nil{
		fmt.Println(">>>>>")
		log.Println(err)
	}
	log.Println(has)
	return services
}
//插入一条记录
func (s *Servicesp) Insert(account,pwd,realname,sid string)(error){
	s.Account=account
	s.Sid=sid
	s.Password=pwd
	s.Realname=realname
	s.Addtime=time.Now().Unix()
	s.Online=1
	fmt.Println(s)
	if affectid,err:=Orm.Table("servicesp").InsertOne(s);err!=nil{
		return err
	}else{
		fmt.Println(affectid)
	}
	return nil
}
