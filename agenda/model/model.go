package model

import (
	"github.com/ZhenlyChen/Agenda-CLI/agenda/util"
	"github.com/spf13/viper"
)

// DBFile 数据库文件路径
type DBFile struct {
	User    string
	Meeting string
	Status  string
}

var DB DBFile

func InitDB() error {
	DB = DBFile{
		User: viper.GetString("DataBase.User"),
		Meeting: viper.GetString("DataBase.Meeting"),
		Status: viper.GetString("DataBase.Status"),
	}
	if err := util.CheckFile(DB.User); err != nil {
		return err
	}
	if err := util.CheckFile(DB.Meeting); err != nil {
		return err
	}
	if err := util.CheckFile(DB.Status); err != nil {
		return err
	}


	return nil
}
