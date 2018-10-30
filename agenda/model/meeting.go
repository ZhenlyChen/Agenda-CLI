package model

import (
	"errors"
)

// MeetingModel 会议数据对象
type MeetingModel struct {
	Data struct {
		Meetings []MeetingData `json:"meetings"`
	} `json:"data"`
	baseModel
}

// MeetingData 会议数据结构
type MeetingData struct {
	Title        string   `json:"title"`
	Presenter    string   `json:"presenter"`
	Participator []string `json:"participator"`
	Start        int64    `json:"start"`
	End          int64    `json:"end"`
}

// GetData 获取数据对象
func (m *MeetingModel) GetData() interface{} {
	return &m.Data
}

// Add 添加一个会议
func (m *MeetingModel) Add(data MeetingData) error {
	m.Data.Meetings = append(m.Data.Meetings, data)
	return m.save(m.Data)
}

// Exist 判断会议是否存在
func (m *MeetingModel) Exist(title string) bool {
	for _, u := range m.Data.Meetings {
		if u.Title == title {
			return true
		}
	}
	return false
}

// 查询指定时间内的会议
func (m *MeetingModel) Query(username string,start int64, end int64) []MeetingData {
	var data []MeetingData
	for _, u := range m.Data.Meetings {
		if start <= u.Start && end >= u.End &&(username == u.Presenter || isParticipator(username, u.Participator)){
			data = append(data, u)
		}
	}
	return data
}

// 是否是参与者
func isParticipator(username string, Participator []string) bool{
	for _, u := range Participator {
		if u == username {
			return true
		}
	}
	return false
}

// 通过title获取会议
func (m *MeetingModel) GetMeetingByTitle(title string) (MeetingData, error) {
	for _, u := range m.Data.Meetings {
		if u.Title == title {
			return u, nil
		}
	}
	return MeetingData{}, errors.New("Not_Meeting")
}

// 通过用户名获取会议
func (m *MeetingModel) GetMeetingByName(name string) []MeetingData {
	var data []MeetingData
	for _, u := range m.Data.Meetings {
		if u.Presenter == name || isParticipator(name, u.Participator) {
			data = append(data, u)
		}
	}
	return data
}

// 获得自己发起的会议
func (m *MeetingModel) GetMeetingAsPresenter(name string) []MeetingData {
	var data []MeetingData
	for _, u := range m.Data.Meetings {
		if u.Presenter == name{
			data = append(data, u)
		}
	}
	return data
}

// 添加参与者
func(m *MeetingModel) AddParticipator(title string, participator []string) error {
	for i, u := range m.Data.Meetings {
		if u.Title == title {
			for _, p := range participator {
				m.Data.Meetings[i].Participator = append(m.Data.Meetings[i].Participator, p)
			}
		}
	}
	return m.save(m.Data)
}

// 是否为会议拥有者
func (m *MeetingModel)IsPresenter(title string, user string) bool {
	meeting, _ := m.GetMeetingByTitle(title)
	return meeting.Presenter == user
}

// 删除会议
func (m *MeetingModel) Delete(title string) error {
	for i, u := range m.Data.Meetings {
		if u.Title == title {
			m.Data.Meetings = append(m.Data.Meetings[:i],m.Data.Meetings[i+1:]...)
			return m.save(m.Data)
		}
	}
	return nil
}