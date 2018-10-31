package controller

import (
	"os"
	"regexp"

	"github.com/ZhenlyChen/Agenda-CLI/agenda/model"
	"github.com/ZhenlyChen/Agenda-CLI/agenda/service"
	"github.com/ZhenlyChen/Agenda-CLI/agenda/util"
)

// UserInterface 用户控制接口
type UserInterface interface {
	Register()
	Login()
	Status()
	Logout()
	List()
	Delete()
	CheckLogin()
}

// Register 注册
func (c *ctrlManger) Register() {
	// 检验参数合法性
	username, err := c.cmd.Flags().GetString("user")
	if err != nil || username == "" {
		util.PrintError("Register Failed! Invalid user name.")
		return
	}

	password, err := c.cmd.Flags().GetString("password")
	if err != nil || password == "" {
		util.PrintError("Register Failed! Invalid password.")
		return
	}

	email, _ := c.cmd.Flags().GetString("email")
	//if email != "" && !regexp.MustCompile("^([A-Za-z0-9_\\-\\.])+\\@([A-Za-z0-9_\\-\\.])+\\.([A-Za-z]{2,8})$").Match([]byte(email)) {
	//	util.PrintError("Register Failed! Invalid email.")
	//	return
	//}

	tel, _ := c.cmd.Flags().GetString("tel")
	if tel != "" && !regexp.MustCompile("^[0-9-+]{3,18}$").Match([]byte(tel)) {
		util.PrintError("Register Failed! Invalid tel.")
		return
	}
	// 调用服务
	err = service.User().Register(model.UserData{
		Name:     username,
		Password: password,
		Email:    email,
		Tel:      tel,
	})
	// 输出结果
	if err == nil {
		util.PrintSuccess("Register Success! Hi, " + username)
	} else if err == service.ErrorRegisterDuplicateName {
		util.PrintError("Register Failed! The user name already exists.")
	} else if err ==  service.ErrorRegisterIllegalName {
		util.PrintError("Register Failed! Illegal user name.")
	} else {
		util.PrintError("Register Failed!")
	}
}

// Login 登陆
func (c *ctrlManger) Login() {
	// 参数合法性
	username, err := c.cmd.Flags().GetString("user")
	if err != nil || username == "" {
		util.PrintError("Login Failed! Invalid user name.")
		return
	}

	password, err := c.cmd.Flags().GetString("password")
	if err != nil || password == "" {
		util.PrintError("Login Failed! Invalid password.")
		return
	}
	// 调用登陆服务
	err = service.User().Login(username, password)
	// 显示结果
	if err == nil {
		util.PrintSuccess("Login Success! Hello, " + username + ".")
	} else if err == service.ErrorLoginNullUser || err == service.ErrorLoginErrorPassword {
		util.PrintError("Login Failed! Incorrect username or password.")
	} else {
		util.PrintError("Login Failed!")
	}
}

// Status 查看登陆状态
func (c *ctrlManger) Status() {
	user := service.Status().GetLoginUser()
	if user == "" {
		util.PrintInfo("No logged in users.")
	} else {
		util.PrintInfo("Current user: " + user + ".")
	}
}

// Logout 退出登陆
func (c *ctrlManger) Logout() {
	service.Status().ClearStatus()
	util.PrintInfo("sign out.")
}

// List 列出所有用户
func (c *ctrlManger) List() {
	users := model.User().GetAllUsers()
	spilt := "    "
	for _, u := range users{
		util.PrintInfo("Username: " + u.Name + spilt + "Email: " + u.Email + spilt + "Tel: " + u.Tel)
	}
}

// Delete 删除当前账户
func (c *ctrlManger) Delete() {
	err := service.User().UserDelete()
	if err == nil {
		util.PrintError("Delete User Success!")
		service.Status().ClearStatus()
	} else{
		util.PrintError("Delete User Failed!")
	}
}

func (c *ctrlManger) CheckLogin(){
	if service.Status().GetLoginUser() == "" {
		util.PrintError("You must login first to obtain permission!")
		os.Exit(0)
	}
}
