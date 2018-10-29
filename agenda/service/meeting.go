package service

import "github.com/ZhenlyChen/Agenda-CLI/agenda/model"

type MeetingInterface interface {
	Create(data model.MeetingData) error
}

func Create(data model.MeetingData) error{

}
