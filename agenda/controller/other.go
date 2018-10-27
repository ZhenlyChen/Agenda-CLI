package controller

import (
	"github.com/ZhenlyChen/Agenda-CLI/agenda/util"
)

// OtherInterface 其他命令接口
type OtherInterface interface {
	Zhenlychen()
	Version()
}

// Zhenlychen 专属命令
func (c *ctrlManger) Zhenlychen() {
	name, err := c.cmd.Flags().GetString("name")
	if err != nil {
		util.PrintError("Invalid name!")
	}
	util.PrintInfo("Your name is " + name)
}

// Version 查看版本信息
func (c *ctrlManger) Version() {
	util.PrintInfo("V1.0.0")
}