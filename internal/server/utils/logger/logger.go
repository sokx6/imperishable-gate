package logger

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

var echoInstance *echo.Echo
var currentLogLevel LogLevel

// LogLevel 定义日志级别
type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	SILENT
)

// InitLogger 初始化日志器，需要传入 Echo 实例
func InitLogger(e *echo.Echo) {
	echoInstance = e

	// 根据环境变量设置日志级别
	currentLogLevel = getLogLevelFromEnv()

	// 同时配置 Echo 的日志级别
	setEchoLogLevel(e)
}

// getLogLevelFromEnv 从环境变量获取日志级别
func getLogLevelFromEnv() LogLevel {
	logLevel := strings.ToLower(os.Getenv("LOG_LEVEL"))
	switch logLevel {
	case "debug":
		return DEBUG
	case "info":
		return INFO
	case "warn", "warning":
		return WARN
	case "error":
		return ERROR
	case "silent":
		return SILENT
	default:
		return WARN // 默认为 WARN
	}
}

// setEchoLogLevel 设置 Echo 框架的日志级别
func setEchoLogLevel(e *echo.Echo) {
	switch currentLogLevel {
	case DEBUG:
		e.Logger.SetLevel(log.DEBUG)
	case INFO:
		e.Logger.SetLevel(log.INFO)
	case WARN:
		e.Logger.SetLevel(log.WARN)
	case ERROR:
		e.Logger.SetLevel(log.ERROR)
	case SILENT:
		e.Logger.SetLevel(log.OFF)
	default:
		e.Logger.SetLevel(log.WARN)
	}
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

// ShouldLogHTTP 判断是否应该记录 HTTP 请求日志（仅在 DEBUG 级别）
func ShouldLogHTTP() bool {
	return currentLogLevel <= DEBUG
}

// Info 记录信息日志
func Info(format string, args ...interface{}) {
	if currentLogLevel > INFO {
		return
	}
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
	if currentLogLevel > INFO {
		return
	}
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
	if currentLogLevel > WARN {
		return
	}
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
	if currentLogLevel > ERROR {
		return
	}
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
	if currentLogLevel > DEBUG {
		return
	}
	message := fmt.Sprintf(format, args...)
	log := formatLog("DEBUG", ColorDebug, message)
	if echoInstance != nil {
		echoInstance.Logger.Debug(log)
	} else {
		fmt.Println(log)
	}
}
