package model

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"shuoba/utils"
)

var Orm *xorm.Engine



func SetEngine() *xorm.Engine{
	username := utils.Cfg.MustValue("db","username","root")
	password := utils.Cfg.MustValue("db","password","abcd@1234")
	server := utils.Cfg.MustValue("db","server","127.0.0.1")
	dbName := utils.Cfg.MustValue("db","db_name","root")
	port := utils.Cfg.MustValue("db","port","3306")
	fmt.Println(username,password,server,dbName,port)
	var err error
	Orm, err := xorm.NewEngine("mysql", username+":"+password+"@("+server+":"+port+")/"+dbName+"?charset=utf8")
	//orm, err := xorm.NewEngine("mysql", "root:123456@(127.0.0.1)/chattalk?charset=utf8")
	if err !=nil{
		log.Println("orm connenction error",err)
	}else{
		log.Println("sql connection  success")
		if err:=Orm.Ping();err!=nil{
			fmt.Println("ping",err)
		}
	}
	Orm.ShowSQL(true)
	return Orm
}
type uitls struct {

}
func init(){
	Orm=SetEngine()
	//fmt.Println(Orm.Ping())
}