package service

import (
	"github.com/ZhenlyChen/Agenda-CLI/agenda/model"
	"github.com/ZhenlyChen/Agenda-CLI/agenda/util"
	"errors"
	"regexp"
)

// UserInterface 用户服务接口
type UserInterface interface {
	Register(data model.UserData) error
	Login(user, password string) error
	UserDelete() error
}

var(
	ErrorRegisterIllegalName = errors.New("illegal_name")
	ErrorRegisterDuplicateName = errors.New("duplicate_name")
	ErrorLoginNullUser = errors.New("null_user")
	ErrorLoginErrorPassword = errors.New("error_password")
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

// Login 登陆
func (s *service) Login(user, password string) error {
	u := s.userModel.GetByName(user)
	// 检查用户
	if u.Name == "" {
		return ErrorLoginNullUser
	}
	// 检查密码
	if u.Password != util.HashPassword(password, u.PasswordSalt) {
		return ErrorLoginErrorPassword
	}
	// 设置登陆状态
	s.SetStatus(u.Name)
	return nil
}

// Delete 删除
func (s *service) UserDelete() error {
	// 删除自己发起的会议
	s.Clear()
	// 退出参与的会议
	meetings := s.meetingModel.GetMeetingAsParticipator(Status().GetLoginUser())
	for _, meeting := range meetings {
		err := s.Quit(meeting.Title)
		if err != nil {
			return err
		}
	}
	// 用户删除
	return s.userModel.Delete(Status().GetLoginUser())
}