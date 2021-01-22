package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

type FileLogger struct {
	Level       LogLevel //日志等级
	fileName    string   // 文件名
	fileObj     *os.File //普通日志级别
	fileerrObj  *os.File //err日志级别
	maxFileSize int64    // 最大切割点
}

func NewFileLogger(levelStr, fn string, maxfs int64) *FileLogger {

	lever, err := ParseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	f1 := &FileLogger{Level: lever, fileName: fn, maxFileSize: maxfs}
	f1.initFile()
	return f1
}

//初始化打开文件生成文件对象
func (F *FileLogger) initFile() {
	fileObj, err := os.OpenFile(F.fileName+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open file failed ,err%v\n", err)
		return
	}
	fileerrObj, err := os.OpenFile(F.fileName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open file failed ,err%v\n", err)
		return
	}
	F.fileObj = fileObj
	F.fileerrObj = fileerrObj
}

/**
判断文件大小
*/
func (F *FileLogger) checkFileSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file failed err:%v\n", err)
		return false
	}
	if fileInfo.Size() >= F.maxFileSize {
		return true
	}
	return false
}

// 切割文件
func (F *FileLogger) splitFile(file *os.File) (*os.File, error) {
	var (
		bakName     string = ""
		NewFileName string = ""
	)
	//判断是否达到文件最大存储内存
	if F.checkFileSize(file) {

		//先获取文件名称再关闭
		fileInfo, err := file.Stat()
		if err != nil {
			return nil, err
		}
		logName := fileInfo.Name()
		//1.关闭文件
		file.Close()
		//2.备份文件
		newStr := time.Now().Format("20060102030405000")
		if logName == path.Base(F.fileName)+".log" {
			bakName = fmt.Sprintf("%s.log.bak%s", F.fileName, newStr)
			NewFileName = F.fileName + ".log"
		} else {
			bakName = fmt.Sprintf("%s.err.bak%s", F.fileName, newStr)
			NewFileName = F.fileName + ".err"
		}
		err = os.Rename(NewFileName, bakName)
		if err != nil {
			panic(err)
		}
		fileObj, err := os.OpenFile(NewFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("open file failed ,err%v\n", err)
			return nil, err
		}
		return fileObj, nil
	}
	return file, nil
}

func (F *FileLogger) log(lev LogLevel, format string, a ...interface{}) {
	if F.Level <= lev {
		msg := fmt.Sprintf(format, a...)
		//获取当前时间
		now := time.Now().Format("2006-01-02 15:04:05")

		//获取文件名称，方法名称，行号
		fileName, funcNmae, line := getLogInfo(3)
		is_lev := []string{UNKNOWN: "UNKNOWN", DEBUG: "DEBUG", TRACE: "TRACE", INFO: "INFO", WARNING: "WARNING", ERROR: "ERROR", FATAL: "FATAL"}
		levStr := is_lev[lev]
		if lev < ERROR {
			fileObj, err := F.splitFile(F.fileObj)
			if err != nil {
				fmt.Printf("init log failed err:%v\n", err)
			}
			F.fileObj = fileObj
			fmt.Fprintf(F.fileObj, "[%s] [%s] [%s:%s:%d]%s\n", now, levStr, fileName, funcNmae, line, msg)
		} else {
			errfileObj, err := F.splitFile(F.fileerrObj)
			if err != nil {
				fmt.Printf("init log failed err:%v\n", err)
			}
			F.fileerrObj = errfileObj
			fmt.Fprintf(F.fileerrObj, "[%s] [%s] [%s:%s:%d]%s\n", now, levStr, fileName, funcNmae, line, msg)
		}
	}
}
func (L *FileLogger) Debug(format string, a ...interface{}) {
	L.log(DEBUG, format, a...)
}
func (L *FileLogger) Info(format string, a ...interface{}) {
	L.log(INFO, format, a...)

}
func (L *FileLogger) Warning(format string, a ...interface{}) {
	L.log(WARNING, format, a...)
}
func (L *FileLogger) Error(format string, a ...interface{}) {
	L.log(ERROR, format, a...)
}
func (L *FileLogger) Fatal(format string, a ...interface{}) {
	L.log(FATAL, format, a...)
}
