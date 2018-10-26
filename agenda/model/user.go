package model

// UserModel 用户数据对象
type UserModel struct {
	Data struct {
		Users []UserData `json:"users"`
	} `json:"data"`
	baseModel
}

// UserData 用户数据结构
type UserData struct {
	Name         string `json:"name"`
	Password     string `json:"password"`
	PasswordSalt string `json:"password_salt"`
	Email        string `json:"email"`
	Tel          string `json:"tel"`
}

// GetData 获取数据对象
func (m *UserModel) GetData() interface{} {
	return &m.Data
}

// Add 添加一个用户
func (m *UserModel) Add(data UserData) error {
	// 操作数据库
	m.Data.Users = append(m.Data.Users, data)
	return m.save(m.Data)
}

// Exist 判断用户是否存在
func (m *UserModel) Exist(name string) bool {
	for _, u := range m.Data.Users {
		if u.Name == name {
			return true
		}
	}
	return false
}