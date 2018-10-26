package service

import (
	"github.com/ZhenlyChen/Agenda-CLI/agenda/model"
	"github.com/ZhenlyChen/Agenda-CLI/agenda/util"
	"github.com/kataras/iris/core/errors"
	"regexp"
)

// UserInterface 用户服务接口
type UserInterface interface {
	Register(data model.UserData) error
}

var(
	ErrorRegisterIllegalName = errors.New("illegal_name")
	ErrorRegisterDuplicateName = errors.New("duplicate_name")
)

// Register 注册
func (s *service) Register(data model.UserData) error {
	// 检查名字合法性
	regName := regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9_-]*$")
	if !regName.Match([]byte(data.Name)) {
		return ErrorRegisterIllegalName
	}
	// 检查重复
	if s.userModel.Exist(data.Name) {
		return ErrorRegisterDuplicateName
	}
	// 加密密码
	data.PasswordSalt = util.GetRandomSalt()
	data.Password = util.HashPassword(data.Password, data.PasswordSalt)
	return s.userModel.Add(data)
}
