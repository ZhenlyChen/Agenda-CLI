package util

import (
	"os"
	"strings"
	"time"
)

var l *logTool

const (
	LogInfo    = "INFO"
	LogWarning = "WARNING"
	LogSuccess = "SUCCESS"
	LogError   = "ERROR"
)

func Log() *logTool { return l }

type logTool struct {
	FilePath string
	UserName string
}

// 初始化
func init() {
	l = new(logTool)
	l.UserName = "Anonymous"
}

func (l *logTool) Init(filePath string) error {
	if err := CheckFile(filePath); err != nil {
		return err
	}
	l.FilePath = filePath
	return nil
}

func (l *logTool) SetUserName(name string) {
	l.UserName = name
}

func (l *logTool) AddLog(logType, content string) {
	logTime := time.Now().Format("2006-01-02 15:04:05")
	l.appendToFile(strings.Join([]string{logTime, "[" + logType + "]", "[" + l.UserName + "]", content, "\n"}, " "))
}

func (l *logTool) appendToFile(text string) {
	fd, _ := os.OpenFile(l.FilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	fd.Write([]byte(text))
	fd.Close()
}