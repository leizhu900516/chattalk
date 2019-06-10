package utils

import (
	cfg "github.com/Unknwon/goconfig"
)

var Cfg *cfg.ConfigFile

func SetSetting(){
	var err error
	Cfg,err = cfg.LoadConfigFile("setting.ini")
	if err!=nil{
		Cfg,err = cfg.LoadConfigFile("../setting.ini")
	}
}

func init() {
	SetSetting()
}