package service

import (
	"errors"
	"github.com/ZhenlyChen/Agenda-CLI/agenda/model"
	"time"
)

type MeetingInterface interface {
	Create(data model.MeetingData, start string, end string) error
}

var(
	ErrorTimeOutOfRange = errors.New("Out_Of_Range")
	ErrorTimeEndTimeEarly = errors.New("End_Time_Early")

)

// 创建会议
func (s *service) Create(data model.MeetingData, start string, end string) error{
	// 检测时间合法性
	startTime, err := time.Parse("2006/01/02-15:04", start)
	if err != nil {
		return ErrorTimeOutOfRange
	}
	endTime, err := time.Parse("2006/01/02-15:04", end)
	if err != nil {
		return ErrorTimeOutOfRange
	}
	if endTime.Unix() <= startTime.Unix(){
		return ErrorTimeEndTimeEarly
	}

	data.Start = startTime.Unix()
	data.End = endTime.Unix()
	s.meetingModel.Add(data)
	return nil
}