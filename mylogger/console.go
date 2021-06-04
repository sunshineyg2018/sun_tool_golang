// 自定义一个日志库
package mylogger

import (
	"fmt"
	"time"
)

// Logger 日志结构体
type ConsoleLogger struct {
	level loglevel
}

// Newlog 构造函数
func Newlog(level string)ConsoleLogger{
	// 判断日志级别
	lervel,err := parseLoglevel(level)
	if err != nil{
		panic(err) // 直接报错
	}
	return ConsoleLogger{
		level: lervel,
	}
}

func (c ConsoleLogger) enable(logLevel loglevel)bool {
	// 如果上面level的值 大于当前的我要打印的值
	return c.level >= logLevel
}



func (c ConsoleLogger)log(lv loglevel,format string,a... interface{}) {
	now := time.Now().Format("2006-01-02 15:04:05")
	funcName, fileName, lineNo := getInfo(3)
	s1 := ""
	switch lv {
	case DEBUG:
		s1 = "DEBUG"
	case FATAL:
		s1 = "FATAL"
	case ERROR:
		s1 = "ERROR"
	case WARNING:
		s1 = "WARNING"
	case INFO:
		s1 = "INFO"
	case TRACE:
		s1 = "TRACE"
	}

	msg := fmt.Sprintf(format, a...) //将字符串格式化为一个字符串
	fmt.Printf("[%s] [%s] [%s:%s:%d] %s \n", now, s1, funcName, fileName, lineNo, msg)
}



func (c ConsoleLogger)Debug(format string,a... interface{}){
	// 判断日志级别再打印
	if c.enable(DEBUG){
		c.log(DEBUG,format,a...)
	}

}

func (c ConsoleLogger)Info(format string,a... interface{}){
	if c.enable(INFO){
		c.log(INFO,format,a...)
	}
}

func (c ConsoleLogger)Warning(format string,a... interface{}){
	if c.enable(WARNING){
		c.log(WARNING,format,a...)
	}
}

func (c ConsoleLogger)Error(format string,a... interface{}){
	if c.enable(ERROR){
		c.log(ERROR,format,a...)
	}
}

func (c ConsoleLogger)Fatal(format string,a... interface{}){
	if c.enable(FATAL){
		c.log(FATAL,format,a...)
	}
}

func (c ConsoleLogger)Close(){

}