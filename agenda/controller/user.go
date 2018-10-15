package controller

import (
	"fmt"
	"os"

	"github.com/ZhenlyChen/Agenda-CLI/agenda/model"
	"github.com/ZhenlyChen/Agenda-CLI/agenda/service"
	"github.com/ZhenlyChen/Agenda-CLI/agenda/util"
)

// UserInterface 用户控制接口
type UserInterface interface {
	Register()
}

// Register 注册
func (c *ctrlManger) Register() {
	// 检验参数合法性
	username, err := c.cmd.Flags().GetString("user")
	if err != nil || username == "" {
		fmt.Fprintln(os.Stderr, "Invalid user name")
		util.Log().AddLog(util.LogError, "Register a user, Failed: Invalid user name")
		return
		// 输出错误信息
	}

	fmt.Println("register called by " + username)
	service.User().Register(model.UserData{
		Name: username,
	})
	// 输出结果

	// 记录日志
	util.Log().AddLog(util.LogSuccess, "User \""+username+"\" register.")
}
