package controller

import (
	"github.com/ZhenlyChen/Agenda-CLI/agenda/model"
	"github.com/ZhenlyChen/Agenda-CLI/agenda/service"
	"github.com/ZhenlyChen/Agenda-CLI/agenda/util"
	"regexp"
	"strings"
	"time"
)

// MeetingInterface 会议控制接口
type MeetingInterface interface {
	Create()
	AddParticipator()
	RemoveParticipator()
	Query()
	MeetingDelete()
	Clear()
}

// Create 创建会议
func (c *ctrlManger) Create(){
	title, err := c.cmd.Flags().GetString("title")
	if err != nil || title == "" {
		util.PrintError("Create Meeting Failed! Invalid title .")
	}

	participator, err := c.cmd.Flags().GetString("participator")
	if err != nil || participator == "" {
		util.PrintError("Create Meeting Failed! Invalid participator .")
	}

	start, err := c.cmd.Flags().GetString("start")
	if err != nil || !regexp.MustCompile("^[0-9]{4}/[0-9]{2}/[0-9]{2}-[0-9]{2}:[0-9]{2}$").Match([]byte(start)) {
		util.PrintError("Create Meeting Failed! Invalid start time .")
	}

	end, err := c.cmd.Flags().GetString("end")
	if err != nil || !regexp.MustCompile("^[0-9]{4}/[0-9]{2}/[0-9]{2}-[0-9]{2}:[0-9]{2}$").Match([]byte(end)) {
		util.PrintError("Create Meeting Failed! Invalid end time .")
	}

	startTime, _ := time.Parse(start, "2006/01/02-15:04")
	endTime, _ := time.Parse(end, "2006/01/02-15:04")

	err := service.Meeting().Create(model.MeetingData{
		Title: title,
		Presenter: service.Status().GetLoginUser(),
		Participator: strings.Split(participator, "+"),
		Start: startTime.Unix(),
		End: endTime.Unix(),
	})
}

// AddParticipator 添加会议参与者
func (c *ctrlManger) AddParticipator(){
	//TODO
}

// RemoveParticipator 删除会议参与者
func (c *ctrlManger) RemoveParticipator(){
	//TODO
}

// Query 会议查询
func (c *ctrlManger) Query(){
	//TODO
}

// MeetingDelete 删除会议
func (c *ctrlManger) MeetingDelete(){
	//TODO
}

// Clear 删除会议
func (c *ctrlManger) Clear(){
	//TODO
}
