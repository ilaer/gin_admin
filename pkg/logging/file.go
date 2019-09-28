/*
@Time : 2019/9/27 上午10:31
@Author : ila
@File : file
@Software: GoLand
*/
package logging
import (
	"fmt"
	"log"
	_ "log"
	"os"
	"time"
)
var(
	LogSavePath="runtime/logs"
	LogSaveName="log"
	LogFileExt="log"
	TimeFormat="20060102"
)
func getLogFilePath() string{
	return fmt.Sprintf("%s",LogSavePath)
}
func getLogFileFullPath() string{
	prefixPath:=getLogFilePath()
	suffixpath:=fmt.Sprintf("%s%s.%s",LogSaveName,time.Now().Format(TimeFormat),LogFileExt)

	return fmt.Sprintf("%s%s",prefixPath,suffixpath)
}

func openLogFile(filePath string) *os.File{
	_,err:=os.Stat(filePath)
	switch{
	case os.IsNotExist(err):
		mkDir()
	case os.IsPermission(err):
		log.Fatalf("Permission:%v",err)
	}
	handle,err:=os.OpenFile(filePath,os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
	if err!=nil{
		log.Fatalf("Fail to OpenFile:%v",err)
	}
	return handle
}

func mkDir(){
	dir,_:=os.Getwd()
	err:=os.MkdirAll(dir+"/"+getLogFilePath(),os.ModePerm)
	if err!=nil{
		panic(err)
	}
}













