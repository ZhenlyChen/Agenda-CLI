package model

// StatusModel 状态数据对象
type StatusModel struct {
	Data StatusData `json:"data"`
	baseModel
}

// StatusData 状态数据结构
type StatusData struct {
	User    string `json:"user"`
	Expires int64  `json:"expires"`
}

// GetData 获取数据对象
func (m *StatusModel) GetData() interface{} {
	return &m.Data
}
