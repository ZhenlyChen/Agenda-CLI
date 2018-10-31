package service

import "github.com/ZhenlyChen/Agenda-CLI/agenda/model"

var s *service

// User 获取用户服务层
func User() UserInterface { return s }

// Meeting 获取会议服务层
func Meeting() MeetingInterface { return s }

// Status 获取状态服务层
func Status() StatusInterface { return s }

type service struct {
	userModel    *model.UserModel
	meetingModel *model.MeetingModel
	statusModel  *model.StatusModel
}

func init() {
	s = new(service)
	s.userModel = model.User()
	s.meetingModel = model.Meeting()
	s.statusModel = model.Status()
}
