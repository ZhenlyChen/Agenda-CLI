package model

import "time"

// StatusModel 状态数据对象
type StatusModel struct {
	Data StatusData `json:"data"`
	baseModel
}

// StatusData 状态数据结构
type StatusData struct {
	User    string `json:"user"` // 登陆用户名
	Expires int64  `json:"expires"` // 设置时间
}

// GetData 获取数据对象
func (m *StatusModel) GetData() interface{} {
	return &m.Data
}

// GetUser 获取当前登陆状态
func (m *StatusModel) GetStatus() StatusData {
	return m.Data
}

// ClearStatus 清除登陆状态
func (m *StatusModel) ClearStatus() {
	m.Data.User = ""
	m.Data.Expires = 0
	m.save(m.Data)
}

// SetUser 设置登陆状态
func (m *StatusModel) SetUser(name string) {
	m.Data.User = name
	m.Data.Expires = time.Now().Unix()
	m.save(m.Data)
}

// RefreshTime 刷新有效时间
func (m *StatusModel) RefreshTime() {
	m.Data.Expires = time.Now().Unix()
	m.save(m.Data)
}