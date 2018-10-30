package service

import (
	"errors"
	"github.com/ZhenlyChen/Agenda-CLI/agenda/model"
	"time"
)

type MeetingInterface interface {
	Create(data model.MeetingData, start string, end string)(string, error)
	AddParticipator(title string, participator []string) error
	RemoveParticipator(title string, participator []string) error
	Query(username string,start string, end string) ([]model.MeetingData, error)
	Delete(title string) error
	Quit(title string) error
	Clear() error
}

var(
	ErrorMeetingNotExist = errors.New("Meeting_Not_Exist")
	ErrorUserNotExist= errors.New("User_Not_Exist")
	ErrorTimeOutOfRange = errors.New("Out_Of_Range")
	ErrorTimeEndTimeEarly = errors.New("End_Time_Early")
	ErrorMeetingDuplicateTitle = errors.New("duplicate_Title")
	ErrorMeetingOverlap = errors.New("Meeting_Overlap")
	ErrorParticipatorExist = errors.New("Participator_Exist")
	ErrorParticipatorNotExist = errors.New("Participator_Not_Exist")
	ErrorNotPresenter= errors.New("Not_Presenter")
	ErrorBothPresenterAndParticipator = errors.New("Both_Presenter_and_Participator")
)

// 检测输入时间合法性
func checkTime(start string, end string) (int64, int64, error) {
	startTime, err := time.Parse("2006/01/02-15:04", start)
	if err != nil {
		return 0, 0, ErrorTimeOutOfRange
	}
	endTime, err := time.Parse("2006/01/02-15:04", end)
	if err != nil {
		return 0, 0, ErrorTimeOutOfRange
	}
	if endTime.Unix() <= startTime.Unix(){
		return 0, 0, ErrorTimeEndTimeEarly
	}
	return startTime.Unix(), endTime.Unix(), nil
}

// 创建会议
func (s *service) Create(data model.MeetingData, start string, end string) (string, error){
	// 检测时间合法性和冲突
	startTime, endTime, err := checkTime(start, end)
	if err != nil {
		return "", err
	}
	if s.meetingModel.Exist(data.Title) {
		return "", ErrorMeetingDuplicateTitle
	}
	data.Start, data.End = startTime, endTime
	// 检测参与者是否存在
	for _, u := range data.Participator{
		if !s.userModel.Exist(u) {
			return u, ErrorUserNotExist
		}
	}
	// 检测创建者是否也为参与者
	for _, p := range data.Participator {
		if p == s.GetLoginUser() {
			return "", ErrorBothPresenterAndParticipator
		}
	}
	// 检测参与者和创建者的时间冲突
	for _, u := range append(data.Participator, Status().GetLoginUser()) {
		meetings := s.meetingModel.GetMeetingByName(u)
		for _, m := range meetings {
			if !(m.End <= data.Start || m.Start >= data.End){
				return u, ErrorMeetingOverlap
			}
		}
	}
	s.meetingModel.Add(data)
	return "", nil
}

// 添加参与者
func (s *service) AddParticipator(title string, participator []string) error{
	// 检测目标会议是否存在并获取
	exist := s.meetingModel.Exist(title)
	if !exist {
		return ErrorMeetingNotExist
	}
	meeting := s.meetingModel.GetMeetingByTitle(title)
	// 检测当前用户是否是会议的拥有者
	if meeting.Presenter != Status().GetLoginUser() {
		return ErrorNotPresenter
	}
	// 检测新增参与者是否已经是参与者
	for _, u := range participator {
		if s.meetingModel.IsParticipator(title, u) {
			return ErrorParticipatorExist
		}
	}
	// 检测会议是否存在重叠
	for _, u := range participator {
		meetings := s.meetingModel.GetMeetingByName(u)
		for _, i := range meetings {
			if !(i.Start >= meeting.End || i.End <= meeting.End){
				return ErrorMeetingOverlap
			}
		}
	}
	return s.meetingModel.AddParticipator(title, participator)
}

// 删除参与者
func (s *service) RemoveParticipator(title string, participator []string) error{
	// 检测目标会议是否存在并获取
	exist := model.Meeting().Exist(title)
	if !exist {
		return ErrorMeetingNotExist
	}
	meeting := s.meetingModel.GetMeetingByTitle(title)
	// 检测是否是会议的拥有者
	if meeting.Presenter != Status().GetLoginUser() {
		return ErrorNotPresenter
	}
	// 检测是否已经是参与者
	for _, u := range participator {
		if !s.meetingModel.IsParticipator(title, u) {
			return ErrorParticipatorNotExist
		}
	}

	for _, p := range participator {
		err := s.meetingModel.RemoveParticipator(title, p)
		if err != nil {
			return err
		}
	}
	return nil
}


// 查询会议
func (s *service) Query(username string, start string, end string) ([]model.MeetingData, error){
	startTime, endTime, err := checkTime(start, end)
	if err != nil {
		return nil, err
	}

	data := s.meetingModel.Query(username, startTime, endTime)

	return data, nil
}

// 通过会议标题删除会议
func (s *service) Delete(title string) error {
	// 检测会议是否存在
	if !model.Meeting().Exist(title) {
		return ErrorMeetingNotExist
	}
	// 检测是否为拥有者
	if !model.Meeting().IsPresenter(title, Status().GetLoginUser()) {
		return ErrorNotPresenter
	}
	return model.Meeting().Delete(title)
}

// 退出会议
func (s *service) Quit(title string) error {
	// 检测会议是否存在
	if !model.Meeting().Exist(title) {
		return ErrorMeetingNotExist
	}
	// 检测是否为参与者
	if !s.meetingModel.IsParticipator(title, Status().GetLoginUser()){
		return ErrorParticipatorNotExist
	}
	return model.Meeting().RemoveParticipator(title, s.GetLoginUser())
}

// 清空会议
func (s *service) Clear() error {
	meetings := s.meetingModel.GetMeetingAsPresenter(Status().GetLoginUser())
	for _, meeting := range meetings {
		err := s.meetingModel.Delete(meeting.Title)
		if err != nil {
			return err
		}
	}
	return nil
}