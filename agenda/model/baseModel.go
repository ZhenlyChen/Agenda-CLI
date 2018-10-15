package model

import (
	"encoding/json"
	"io/ioutil"
)

// BaseModel 基本数据
type baseModel struct {
	file string
	impl BaseModel
}
// BaseModel 基本接口
type BaseModel interface {
	GetData() interface{}
}

func (m *baseModel) GetData() interface{} {
	return m.impl.GetData()
}

func (m *baseModel) save(data interface{}) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(m.file, b,0777)
}


