package logger

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
)

var echoInstance *echo.Echo

// InitLogger 初始化日志器，需要传入 Echo 实例
func InitLogger(e *echo.Echo) {
	echoInstance = e
}

// 日志级别颜色
const (
	ColorReset   = "\033[0m"
	ColorInfo    = "\033[36m" // 青色
	ColorSuccess = "\033[32m" // 绿色
	ColorWarning = "\033[33m" // 黄色
	ColorError   = "\033[31m" // 红色
	ColorDebug   = "\033[35m" // 紫色
)

// formatLog 格式化日志消息
func formatLog(level, levelColor, message string) string {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	return fmt.Sprintf(
		"\033[90m%s\033[0m | %s%-8s%s | %s",
		timestamp,
		levelColor,
		level,
		ColorReset,
		message,
	)
}

// Info 记录信息日志
func Info(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	log := formatLog("INFO", ColorInfo, message)
	if echoInstance != nil {
		echoInstance.Logger.Info(log)
	} else {
		fmt.Println(log)
	}
}

// Success 记录成功日志
func Success(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	log := formatLog("SUCCESS", ColorSuccess, message)
	if echoInstance != nil {
		echoInstance.Logger.Info(log)
	} else {
		fmt.Println(log)
	}
}

// Warning 记录警告日志
func Warning(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	log := formatLog("WARNING", ColorWarning, message)
	if echoInstance != nil {
		echoInstance.Logger.Warn(log)
	} else {
		fmt.Println(log)
	}
}

// Error 记录错误日志
func Error(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	log := formatLog("ERROR", ColorError, message)
	if echoInstance != nil {
		echoInstance.Logger.Error(log)
	} else {
		fmt.Println(log)
	}
}

// Debug 记录调试日志
func Debug(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	log := formatLog("DEBUG", ColorDebug, message)
	if echoInstance != nil {
		echoInstance.Logger.Debug(log)
	} else {
		fmt.Println(log)
	}
}
