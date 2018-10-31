package model

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"

	"github.com/ZhenlyChen/Agenda-CLI/agenda/util"
)

// DataFile 数据库文件
type DataFile struct {
	User    string
	Meeting string
	Status  string
}

var m *Model

func init() {
	m = new(Model)
}

// Model 数据管理
type Model struct {
	User    UserModel
	Meeting MeetingModel
	Status  StatusModel
}

// User 获取用户数据操作对象
func User() *UserModel { return &m.User }

// Meeting 获取会议数据操作对象
func Meeting() *MeetingModel { return &m.Meeting }

// Status 获取状态数据操作对象
func Status() *StatusModel { return &m.Status }

// InitDB 初始化数据库
func InitDB(file DataFile) error { return m.initDB(file) }
func (m *Model) initDB(file DataFile) error {
	m.User.baseModel = baseModel{impl: &m.User}
	if err := m.initModel(file.User, &m.User.baseModel); err != nil {
		return err
	}
	m.Meeting.baseModel = baseModel{impl: &m.Meeting}
	if err := m.initModel(file.Meeting, &m.Meeting.baseModel); err != nil {
		return err
	}
	m.Status.baseModel = baseModel{impl: &m.Status}
	if err := m.initModel(file.Status, &m.Status.baseModel); err != nil {
		return err
	}
	return nil
}

func (m *Model) initModel(filePath string, dataModel interface{}) error {
	d, ok := dataModel.(*baseModel)
	if ok != true {
		return errors.New("error type")
	}
	d.file = filePath
	if err := util.CheckFile(filePath); err != nil {
		return err
	}
	// 读取数据
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err // 解析原始数据错误
	}
	if len(data) == 0 {
		return nil // 无初始数据
	}
	if err := json.Unmarshal(data, d.GetData()); err != nil {
		if err != io.EOF {
			return err
		}
	}
	return nil
}

//Used for testing  --ctp
func ClearModel(){
	User().Data.Users = []UserData{}
	User().save(User().Data)
	Status().Data =  StatusData{}
	Status().save(Status().Data )
	Meeting().Data.Meetings = []MeetingData{}
	Meeting().save(Meeting().Data)
}
