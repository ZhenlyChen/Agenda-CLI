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
	MeetingQuit()
	Clear()
}

// Create 创建会议
func (c *ctrlManger) Create(){
	// 获取参数
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

	// 调用service服务
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
	}else if err == service.ErrorMeetingDuplicateTitle {
		util.PrintError("Create Meeting Failed! The Meeting title already exists.")
	}else {
		util.PrintError("Create Meeting Failed!")
	}

}

// AddParticipator 添加会议参与者
func (c *ctrlManger) AddParticipator(){
	// 获取参数
	title, err := c.cmd.Flags().GetString("title")
	if err != nil || title == "" {
		util.PrintError("Add Participator Failed! Invalid title .")
		return
	}

	participator, err := c.cmd.Flags().GetString("participator")
	if err != nil || participator == "" {
		util.PrintError("Add Participator Failed! Invalid participator .")
		return
	}

	// 调用service服务
	err = service.Meeting().AddParticipator(title,  strings.Split(participator, "+"))
	if err == nil{
		util.PrintSuccess("Create Meeting [" + title + "] Success! .")
	} else if err == service.ErrorMeetingNotExist{
		util.PrintError("Add Participator Failed! Not such Meeting [" + title + "] .")
	} else if err == service.ErrorParticipatorExist{
		util.PrintError("Add Participator Failed! Some participators already exist .")
	} else if err == service.ErrorNotPresenter {
		util.PrintError("Add Participator Failed! You are not the presenter of the meeting .")
	} else {
		util.PrintError("Add Participator Failed!")
	}
}

// RemoveParticipator 删除会议参与者
func (c *ctrlManger) RemoveParticipator(){
	//TODO
}

// Query 会议查询
func (c *ctrlManger) Query(){
	start, err := c.cmd.Flags().GetString("start")
	if err != nil || !regexp.MustCompile("^[0-9]{4}/[0-9]{2}/[0-9]{2}-[0-9]{2}:[0-9]{2}$").Match([]byte(start)) {
		util.PrintError("Query Meeting Failed! Invalid start time .")
		return
	}

	end, err := c.cmd.Flags().GetString("end")
	if err != nil || !regexp.MustCompile("^[0-9]{4}/[0-9]{2}/[0-9]{2}-[0-9]{2}:[0-9]{2}$").Match([]byte(end)) {
		util.PrintError("Query Meeting Failed! Invalid end time .")
		return
	}
	// 调用service服务
	data, err := service.Meeting().Query(service.Status().GetLoginUser(),start, end)
	if err == nil{
		util.PrintError("Query Meeting Success!")
		spilt := "    "
		for _, u := range data{
			util.PrintError("Title: " + u.Title + spilt + "Start: " + time.Unix(u.Start, 0).String() + spilt + "End: " + time.Unix(u.End, 0).String())
		}
	} else if err == service.ErrorTimeOutOfRange{
		util.PrintError("Query Meeting Failed! Start or End time out of range .")
	} else if err == service.ErrorTimeEndTimeEarly{
		util.PrintError("Query Meeting Failed! End time is earlier than Start time .")
	} else {
		util.PrintError("Query Meeting Failed!")
	}

}

// MeetingDelete 删除会议
func (c *ctrlManger) MeetingDelete(){
	//TODO
}

// MeetingQuit 退出会议
func (c *ctrlManger) MeetingQuit(){
	//TODO
}

// Clear 删除会议
func (c *ctrlManger) Clear(){
	//TODO
}
