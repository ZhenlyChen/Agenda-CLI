package util

import (
	"os"
	"path"
)

// 检查文件是否存在并且如果不存在则创建并打开
func CheckFile(filePath string) error {
	dirPath := path.Dir(filePath)
	// 判断文件夹是否存在
	if _, err := os.Stat(dirPath); err != nil && os.IsNotExist(err) {
		if err := os.MkdirAll(dirPath, 0777); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}
	// 判断文件是否存在
	if _, err := os.Stat(filePath); err != nil && os.IsNotExist(err) {
		if _, err := os.Create(filePath); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}
	return nil
}