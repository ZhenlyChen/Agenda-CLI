package service

import "github.com/ZhenlyChen/Agenda-CLI/agenda/model"

// UserInterface 用户服务接口
type UserInterface interface {
	Register(data model.UserData) error
}

// Register 注册
func (s *service) Register(data model.UserData) error {
	// 检查名字合法性
	// 检查重复
	// 加密密码
	return s.userModel.Add(data)
}
