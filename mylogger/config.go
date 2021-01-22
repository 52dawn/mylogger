package mylogger

import (
	"errors"
	"strings"
)

//自定义类型
type LogLevel uint16

//定义日志级别常量
const (
	UNKNOWN LogLevel = iota
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

//日志结构体
type Logger struct {
	Level LogLevel
}

func ParseLogLevel(s string) (LogLevel, error) {

	// 字符串转换为小写
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("无效的日志级别")
		return UNKNOWN, err
	}
}
