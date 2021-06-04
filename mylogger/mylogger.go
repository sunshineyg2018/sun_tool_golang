package mylogger

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
)

// 往终端写日志


type loglevel uint16


//定义接口，loger 允许有下列方法
type Loger interface {
	Debug(format string,a... interface{})
	Info(format string,a... interface{})
	Warning(format string,a... interface{})
	Error(format string,a... interface{})
	Fatal(format string,a... interface{})
	Close()
}


const (
	// 定义日志级别类型
	UNKNOWN loglevel = iota
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)



func parseLoglevel(s string)(loglevel,error){
	s = strings.ToLower(s)   //把字母全部弄为小写
	switch s {
	case "debug":
		return DEBUG,nil
	case "trace":
		return TRACE,nil
	case "info":
		return INFO,nil
	case "warning":
		return WARNING,nil
	case "error":
		return ERROR,nil
	case "fatal":
		return FATAL,nil
	default:
		err := errors.New("无效的错误级别")
		return UNKNOWN,err
	}
}



func getInfo(skip int)(funcName ,fileName string,lineNo int){
	//记录程序运行堆栈信息
	pc,file,line,ok := runtime.Caller(skip)
	if !ok{
		fmt.Println("调用失败")
	}
	funcName = runtime.FuncForPC(pc).Name()

	return funcName,path.Base(file),line

}
