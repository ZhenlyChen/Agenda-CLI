package controller

import (
	"github.com/ZhenlyChen/Agenda-CLI/agenda/model"
	"github.com/ZhenlyChen/Agenda-CLI/agenda/service"
	"github.com/ZhenlyChen/Agenda-CLI/agenda/util"
	"regexp"
	"strings"
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
		return
	}

	participator, err := c.cmd.Flags().GetString("participator")
	if err != nil || participator == "" {
		util.PrintError("Create Meeting Failed! Invalid participator .")
		return
	}

	start, err := c.cmd.Flags().GetString("start")
	if err != nil || !regexp.MustCompile("^[0-9]{4}/[0-9]{2}/[0-9]{2}-[0-9]{2}:[0-9]{2}$").Match([]byte(start)) {
		util.PrintError("Create Meeting Failed! Invalid start time .")
		return
	}

	end, err := c.cmd.Flags().GetString("end")
	if err != nil || !regexp.MustCompile("^[0-9]{4}/[0-9]{2}/[0-9]{2}-[0-9]{2}:[0-9]{2}$").Match([]byte(end)) {
		util.PrintError("Create Meeting Failed! Invalid end time .")
		return
	}
	err = service.Meeting().Create(model.MeetingData{
				Title: title,
				Presenter: service.Status().GetLoginUser(),
				Participator: strings.Split(participator, "+"),
			}, start, end)
	if err == nil{
		util.PrintSuccess("Create Meeting [" + title + "] Success! .")
	}else if err == service.ErrorTimeOutOfRange{
		util.PrintError("Create Meeting Failed! Start or End time out of range .")
	}else if err == service.ErrorTimeEndTimeEarly{
		util.PrintError("Create Meeting Failed! End time is earlier than Start time .")
	}else {
		util.PrintError("Create Meeting Failed!")
	}

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
