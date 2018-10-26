package service

import (
	"github.com/ZhenlyChen/Agenda-CLI/agenda/util"
	"time"
)

// UserInterface 用户服务接口
type StatusInterface interface {
	GetLoginUser() string
	SetUser(name string)
	ClearStatus()
}

// GetLoginUser 获取已登陆的用户名，若没有则为“”
func (s *service) GetLoginUser() string {
	status := s.statusModel.GetStatus()

	if status.User == "" {
		return ""
	}
	// 检查用户名是否有效
	if !s.userModel.Exist(status.User) {
		// 非法用户名
		s.ClearStatus()
		return ""
	}
	// 检查登陆状态是否有效
	expires := time.Unix(status.Expires, 0)
	if time.Now().Sub(expires) > time.Hour * 3 {
		// 登陆状态已失效
		s.ClearStatus()
		return ""
	} else {
		s.statusModel.RefreshTime()
		return status.User
	}
}

// SetUser 设置已登录用户
func (s *service) SetUser(name string) {
	s.statusModel.SetUser(name)
	util.Log().SetUserName(name)
}

// ClearStatus 清除登陆状态
func (s *service) ClearStatus() {
	s.statusModel.ClearStatus()
	util.Log().SetUserName("")
}