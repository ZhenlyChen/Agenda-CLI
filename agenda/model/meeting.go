package model

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

/* ------ 查询操作 ------*/

// 获取数据对象
func (m *MeetingModel) GetData() interface{} {
	return &m.Data
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
		if start <= u.Start && end >= u.End &&(username == u.Presenter || m.IsParticipator(u.Title, username)){
			data = append(data, u)
		}
	}
	return data
}

// 是否是参与者
func (m *MeetingModel) IsParticipator(title string, username string) bool{
	meeting := m.GetMeetingByTitle(title)
	for _, p := range meeting.Participator {
		if p == username {
			return true
		}
	}
	return false
}

// 是否为拥有者
func (m *MeetingModel)IsPresenter(title string, user string) bool {
	meeting := m.GetMeetingByTitle(title)
	return meeting.Presenter == user
}

// 通过title获取会议
func (m *MeetingModel) GetMeetingByTitle(title string) (MeetingData) {
	for _, u := range m.Data.Meetings {
		if u.Title == title {
			return u
		}
	}
	return MeetingData{}
}

// 获得作为发起者的会议
func (m *MeetingModel) GetMeetingAsPresenter(name string) []MeetingData {
	var data []MeetingData
	for _, u := range m.Data.Meetings {
		if u.Presenter == name{
			data = append(data, u)
		}
	}
	return data
}

// 获取作为参与者的会议
func (m *MeetingModel) GetMeetingAsParticipator(name string) []MeetingData {
	var data []MeetingData
	for _, u := range m.Data.Meetings {
		for _, p := range u.Participator {
			if p == name{
				data = append(data, u)
				break
			}
		}
	}
	return data
}

// 获取作为发起者或者参与者的会议
func (m *MeetingModel) GetMeetingByName(name string) []MeetingData {
	meetings := m.GetMeetingAsPresenter(name)
	return append(meetings, m.GetMeetingAsParticipator(name)...)
}

/* ------ 增删操作 ------*/

// 添加一个会议
func (m *MeetingModel) Add(data MeetingData) error {
	m.Data.Meetings = append(m.Data.Meetings, data)
	return m.save(m.Data)
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

// 删除参与者
func(m *MeetingModel) RemoveParticipator(title string, participator string) error {
	for i, u := range m.Data.Meetings {
		if u.Title == title {
			for index, p := range u.Participator {
				if p == participator {
					m.Data.Meetings[i].Participator = append(
						m.Data.Meetings[i].Participator[:index],
						m.Data.Meetings[i].Participator[index+1:]...)
				}
			}
			// 删除后无参与者删除会议
			if len(m.Data.Meetings[i].Participator) == 0 {
				err := m.Delete(title)
				if err != nil {
					return err
				}
			}
			break
		}
	}
	return m.save(m.Data)
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