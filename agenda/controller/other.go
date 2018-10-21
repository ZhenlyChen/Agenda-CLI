package controller

import (
	"github.com/ZhenlyChen/Agenda-CLI/agenda/util"
)

// OtherInterface 其他命令接口
type OtherInterface interface {
	Zhenlychen()
}

// Zhenlychen 专属命令
func (c *ctrlManger) Zhenlychen() {
	name, err := c.cmd.Flags().GetString("name")
	if err != nil {
		util.PrintError("Invalid name!")
	}
	util.PrintInfo("Your name is " + name)
}