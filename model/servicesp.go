package model

import "log"

//客服模型

type Servicep struct {
	Id          int 		`form:"id" xorm:"int(11) pk not null autoincr"`
	Name 		string 		`form:"name" xorm:"varchar(50) not null"`
	Sid 		string		`form:"sid" xorm:"varchar(20) not null"`
	Password 	string 		`form:"password" xorm:"varchar(20) not null"`
}

var services []Servicep

func (s *Servicep) getOnlineS() []Servicep{

	has ,err := orm.Table("servicep").Where("online = ?",1).Get(&services)
	if err !=nil{
		log.Println(err)
	}
	log.Println(has)
	return services
}