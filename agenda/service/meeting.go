package service

import (
	"errors"
	"github.com/ZhenlyChen/Agenda-CLI/agenda/model"
	"time"
)

type MeetingInterface interface {
	Create(data model.MeetingData, start string, end string) error
	AddParticipator(title string, participator []string) error
	Query(username string,start string, end string) ([]model.MeetingData, error)
}

var(
	ErrorMeetingNotExist = errors.New("Meeting_Not_Exist")
	ErrorTimeOutOfRange = errors.New("Out_Of_Range")
	ErrorTimeEndTimeEarly = errors.New("End_Time_Early")
	ErrorMeetingDuplicateTitle = errors.New("duplicate_Title")
	ErrorMeetingOverlap = errors.New("Meeting_Overlap")
	ErrorParticipatorExist = errors.New("Participator_Exist")
	ErrorNotPresenter= errors.New("Not_Presenter")
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
func (s *service) Create(data model.MeetingData, start string, end string) error{
	startTime, endTime, err := checkTime(start, end)
	if err != nil {
		return err
	}
	if s.meetingModel.Exist(data.Title) {
		return ErrorMeetingDuplicateTitle
	}

	data.Start = startTime
	data.End = endTime
	s.meetingModel.Add(data)
	return nil
}

// 添加参与者
func (s *service) AddParticipator(title string, participator []string) error{
	// 检测目标会议是否存在
	meeting, err := s.meetingModel.GetMeetingByTitle(title)
	if err != nil {
		return ErrorMeetingNotExist
	}
	// 检测是否是会议的拥有者
	if meeting.Presenter != Status().GetLoginUser() {
		return ErrorNotPresenter
	}
	// 检测是否已经是参与者或拥有者
	for _, u := range participator {
		for _, p := range meeting.Participator {
			if u == p || u == meeting.Presenter {
				return ErrorParticipatorExist
			}
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

// 查询会议
func (s *service) Query(username string, start string, end string) ([]model.MeetingData, error){
	startTime, endTime, err := checkTime(start, end)
	if err != nil {
		return nil, err
	}

	data := s.meetingModel.Query(username, startTime, endTime)

	return data, nil
}