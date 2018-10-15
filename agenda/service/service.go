package service

import "github.com/ZhenlyChen/Agenda-CLI/agenda/model"

var s *service

// User 获取用户服务层
func User() UserInterface { return s }

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
