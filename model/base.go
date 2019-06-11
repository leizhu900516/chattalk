package model

import (
	"github.com/go-xorm/xorm"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"../utils"
)

var orm *xorm.Engine



func SetEngine() *xorm.Engine{
	username := utils.Cfg.MustValue("db","username","root")
	password := utils.Cfg.MustValue("db","username","abcd@1234")
	server := utils.Cfg.MustValue("db","server","127.0.0.1")
	dbName := utils.Cfg.MustValue("db","db_name","root")
	port := utils.Cfg.MustValue("db","port","3306")
	var err error
	orm, err := xorm.NewEngine("mysql", username+":"+password+"@tcp("+server+":"+port+")/"+dbName+"?charset=utf8")
	if err !=nil{
		log.Println("orm connenction error",err)
	}
	orm.ShowSQL(true)
	return orm
}

func init(){
	SetEngine()
}