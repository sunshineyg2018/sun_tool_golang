package main

import (
	"learning/mylogger"
)

// 测试自己写的日志库

//全局结构体
var log mylogger.Loger

func main()  {
	//打印在终端上面
	//log_ := mylogger.Newlog("debug")
	//log_.Debug("记录事件 %d %d",10,2)

	//打印在文件里面 )

	mylogger.NewFileLogger("debug","./","xiaoyang.log",10*1024*1024).Debug("记录事件 %d %d",10,2)
	//log = mylogger.Newlog("debug")
	//log.Debug("这是一条提醒 %d:%s",1001,"杨光")
	//log.Close()

}