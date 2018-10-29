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

// GetData 获取数据对象
func (m *MeetingModel) GetData() interface{} {
	return &m.Data
}

// Add 添加一个会议
func (m *MeetingModel) Add(data MeetingData) interface{} {
	m.Data.Meetings = append(m.Data.Meetings, data)
	return &m.Data
}