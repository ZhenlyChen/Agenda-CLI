package util

import (
	"fmt"
	"os"
)

// PrintError 输出错误并记录日志
func PrintError(content string) {
	fmt.Fprintln(os.Stderr, content)
	Log().AddLog(LogError, content)
}

// PrintSuccess 输出成功并记录日志
func PrintSuccess(content string) {
	fmt.Println(content)
	Log().AddLog(LogSuccess, content)
}

// PrintWarning 输出警告并记录日志
func PrintWarning(content string) {
	fmt.Fprintln(os.Stderr, content)
	Log().AddLog(LogWarning, content)
}

// PrintInfo 输出信息并记录日志
func PrintInfo(content string) {
	fmt.Println(content)
	Log().AddLog(LogInfo, content)
}