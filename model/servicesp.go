package model

import "log"

//客服模型

type Servicep struct {
	Id          int 		`form:"id" xorm:"int(11) pk notnull autoincr"`
	Name 		string 		`form:"name" xorm:"varchar(50) not null"`
	Sid 		string		`form:"sid" xorm:"varchar(20) not null"`
	Addtime 	int 		`form:"addtime" xorm:"int(11) not null"`
	Password 	string 		`form:"password" xorm:"varchar(20) not null"`
	Online 		int			`form:"online" xorm:"int(1) not null default 1"`
}


//获取在线的客服列表提供给轮询算法
func (s *Servicep) getOnlineS() []Servicep{
	var services []Servicep
	has ,err := orm.Table("servicep").Where("online = ?",1).Get(&services)
	if err !=nil{
		log.Println(err)
	}
	log.Println(has)
	return services
}