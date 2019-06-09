package model

import (
	"github.com/go-xorm/xorm"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

var orm *xorm.Engine



func SetEngine() *xorm.Engine{
	username :=""
	password := ""
	server := ""
	dbName := ""
	var err error
	orm, err := xorm.NewEngine("mysql", username+":"+password+"@tcp("+server+":3306)/"+dbName+"?charset=utf8")
	if err !=nil{
		log.Println("orm connenction error",err)
	}
	orm.ShowSQL(true)
	return orm
}