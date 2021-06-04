package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

// 往文件里面写日志相关的信息

type FileLogger struct {
	level loglevel  // 级别
	filePath string //日志文件保存的路径
	fileName string // 日志文件保存的文件名
	maxFileSize int64 // 文件大小
	fileObj *os.File
	errFileObj *os.File
}


func NewFileLogger (levelStr,fp,fn string, maxSize int64)*FileLogger{
	//构造函数，因为很大 所有使用指针
	// 变成loglevel 类型
	logLver, err := parseLoglevel(levelStr)
	if err != nil {
		panic(err)
	}
	// 初始化
	fl := &FileLogger{
		level:       logLver,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
	}
	err = fl.initFile() // 案子文件路径和文件名将文件打开
	if err != nil{
		panic(err)
	}
	return fl
}

func (f *FileLogger)initFile()(error){
	fullFileName := path.Join(f.filePath,f.fileName)
	fileObj,err := os.OpenFile(fullFileName,os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
	if err != nil{
		fmt.Println("open log file failed,err:%v\n",err)
		return err
	}

	errfileObj,err := os.OpenFile(fullFileName+".err",os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
	if err != nil{
		fmt.Println("open log errfile failed,err:%v\n",err)
		return err
	}
	f.fileObj = fileObj
	f.errFileObj = errfileObj
	return nil

}

func (f *FileLogger)Close(){
	f.fileObj.Close()
	f.errFileObj.Close()
}

func (f *FileLogger) enable(logLevel loglevel)bool {
	// 如果上面level的值 大于当前的我要打印的值
	return logLevel >= f.level
}


func (f *FileLogger)checkSize(file *os.File)bool{
	//文件对象的类型
	//根据文件对象的obj 获取文件对象的详细信息
	fileIofo,err := file.Stat()
	if err != nil{
		fmt.Printf("get file info failed,err:%v\n",err)
		return false
	}
	//如果当前文件大小 大于等于 日志文件的最大值 就应该返回true
	return fileIofo.Size() >= f.maxFileSize

}

func (f *FileLogger)splitFile(file *os.File)(*os.File,error)  {
	// 需要切割日志文件
	// 1.备份一下 rename  xx.log -> xx.log.bak201908031709
	nowStr := time.Now().Format("2006010215040500")
	fileInfo,err := file.Stat()
	if err != nil{
		fmt.Printf("get file info failed,err:%v\n",err)
		return nil, err
	}

	logName := path.Join(f.filePath,fileInfo.Name())
	newlogName := fmt.Sprintf("%s.bak%s%s",f.filePath,f.fileName,nowStr)
	// 2.关闭当前的日志文件
	file.Close()
	os.Rename(logName,newlogName)  //改名字了
	// 3.打开一个新的日志文件
	fileObj,err := os.OpenFile(logName,os.O_CREATE|os.O_APPEND|os.O_WRONLY,0644)
	if err != nil{
		fmt.Printf("重新打开新文件发生错误,err:%v",err)
	}
	// 4.将打开的新日志文件对象赋值给 f.fileObj
	return fileObj,err
}

func (f *FileLogger)log(lv loglevel,format string,a... interface{}) {
	now := time.Now().Format("2006-01-02 15:04:05")
	//获取打印行数的数据
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

	if f.checkSize(f.fileObj){
		newFile,err := f.splitFile(f.fileObj)
		if err != nil{
			return
		}
		f.fileObj = newFile
	}
	fmt.Fprintf(f.fileObj,"[%s] [%s] [%s:%s:%d] %s \n", now, s1, funcName, fileName, lineNo, msg)
	if lv >= ERROR{
		newFile,err := f.splitFile(f.errFileObj)
		if err != nil{
			return
		}
		f.errFileObj = newFile
		//如果要记录的日志大于等于ERROR级别,我还要在err记录日志
		fmt.Fprintf(f.errFileObj,"[%s] [%s] [%s:%s:%d] %s \n", now, s1, funcName, fileName, lineNo, msg)
	}
}





func (f *FileLogger)Debug(format string,a... interface{}){
	// 判断日志级别再打印
	if f.enable(DEBUG){
		f.log(DEBUG,format,a...)
	}

}

func (f *FileLogger)Info(format string,a... interface{}){
	if f.enable(INFO){
		f.log(INFO,format,a...)
	}
}

func (f *FileLogger)Warning(format string,a... interface{}){
	if f.enable(WARNING){
		f.log(WARNING,format,a...)
	}
}

func (f *FileLogger)Error(format string,a... interface{}){
	if f.enable(ERROR){
		f.log(ERROR,format,a...)
	}
}

func (f *FileLogger)Fatal(format string,a... interface{}){
	if f.enable(FATAL){
		f.log(FATAL,format,a...)
	}
}
