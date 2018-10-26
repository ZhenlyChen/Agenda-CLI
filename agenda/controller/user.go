package controller

import (
	"regexp"

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
	if  err != nil || username == "" {
		util.PrintError("Register Failed! Invalid user name")
		return
	}

	password, err := c.cmd.Flags().GetString("password")
	if  err != nil || password == "" {
		util.PrintError("Register Failed! Invalid password")
		return
	}

	email, _ := c.cmd.Flags().GetString("email")
	if  email != "" && !regexp.MustCompile("^([A-Za-z0-9_\\-\\.])+\\@([A-Za-z0-9_\\-\\.])+\\.([A-Za-z]{2,8})$").Match([]byte(email)) {
		util.PrintError("Register Failed! Invalid email")
		return
	}

	tel, _ := c.cmd.Flags().GetString("tel")
	if  tel != "" && !regexp.MustCompile("^[0-9-+]{3,18}$").Match([]byte(tel)) {
		util.PrintError("Register Failed! Invalid tel")
		return
	}
	// 调用服务
	err = service.User().Register(model.UserData{
		Name: username,
		Password: password,
		Email: email,
		Tel: tel,
	})
	// 输出结果
	if err == nil {
		util.PrintSuccess("Register Success! Hi, " + username)
	} else if service.ErrorRegisterDuplicateName.Equal(err) {
		util.PrintError("Register Failed! The user name already exists.")
	} else if service.ErrorRegisterIllegalName.Equal(err) {
		util.PrintError("Register Failed! Illegal user name.")
	}else {
		util.PrintError("Register Failed!")
	}
}
