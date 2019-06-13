package model

import (
	"log"
	"time"
)

//聊天消息模型

type Messages struct {
	Mid 		int 		`form:"mid" xorm:"int(11) notnull pk  autoincr"`
	Msg 		string 		`form:"msg" xorm:"text notnull"`		//消息内容
	Addtime 	int 		`form:"addtime" xorm:"int(11)"`
	Updatetime 	int 		`form:"updatetime" xorm:"int(11)"`		 //更新时间
	Unread		int			`form:"unread" xorm:"int(1) default 0"`  //是否已读 0未读1已读
	Userid 		string 		`form:"userid" xorm:"varchar(20) notnull"` //发送人
	Destid 		string 		`form:"destid" xorm:"varchar(20) notnull"` //接收人
}


//获取历史消息，分页，只有用户点击【查看更多】时才会获取
// 每次获取10条
func (m *Messages) GetHistoryMsg(limit,step int)([]Messages,error){
	var data []Messages
	if err := Orm.Table("messages").Limit(limit,limit*(step-1)).Desc("addtime").Find(&data);err!=nil{
		return nil,err
	}
	return data,nil
}
//插入信息
func (m *Messages) Insert() error{
	msg :=new(Messages)
	if affected,err :=Orm.Table("messages").Insert(msg);err!=nil{
		return err
	}else {
		log.Println("插入成功",affected)
	}
	return nil
}
//更新消息状态
//更新unread 和updatetime
func (m *Messages) Update(userid,destid string) error{
	updatetime := time.Now().Unix()
	if affected,err := Orm.Exec("update messages set unread=1,updatetime= ? where userid = ? and destid = ?",updatetime,userid,destid);err != nil{
		log.Println("update error",err)
		return err
	}else{
		log.Println("update success",affected)
	}
	return nil
}