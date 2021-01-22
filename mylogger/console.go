package mylogger

import (
	"fmt"
	"path"
	"runtime"
	"strings"
	"time"
)

//构造函数
func Newlog(levelStr string) Logger {
	lever, err := ParseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return Logger{lever}
}

func getLogInfo(skip int) (fileName, funcName string, lineNo int) {
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		fmt.Printf("err \n")
		return
	}

	//获取方法名
	funcName = runtime.FuncForPC(pc).Name()

	//获取文件名
	fileName = path.Base(file)
	funcName = strings.Split(funcName, ".")[1]
	return
}

//记录日志输出
func (L *Logger) log(lev LogLevel, format string, a ...interface{}) {
	if L.Level <= lev {
		msg := fmt.Sprintf(format, a...)
		//获取当前时间
		now := time.Now().Format("2006-01-02 15:04:05")

		//获取文件名称，方法名称，行号
		fileName, funcNmae, line := getLogInfo(3)
		is_lev := []string{UNKNOWN: "UNKNOWN", DEBUG: "DEBUG", TRACE: "TRACE", INFO: "INFO", WARNING: "WARNING", ERROR: "ERROR", FATAL: "FATAL"}
		levStr := is_lev[lev]
		fmt.Printf("[%s] [%s] [%s:%s:%d]%s\n", now, levStr, fileName, funcNmae, line, msg)
	}
}

func (L *Logger) Debug(format string, a ...interface{}) {
	L.log(DEBUG, format, a...)
}
func (L *Logger) Info(format string, a ...interface{}) {
	L.log(INFO, format, a...)
}
func (L *Logger) Warning(format string, a ...interface{}) {
	L.log(WARNING, format, a...)
}
func (L *Logger) Error(format string, a ...interface{}) {
	L.log(ERROR, format, a...)
}
func (L *Logger) Fatal(format string, a ...interface{}) {
	L.log(FATAL, format, a...)
}
