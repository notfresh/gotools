package gotools

import (
	"encoding/json"
	"fmt"
	"os"
)

type JsonLog struct {
	logger *json.Encoder
	file   *os.File
}

func NewJsonLog(logFileName string) *JsonLog {
	_, err := os.Stat(logFileName)
	flag := err == nil || os.IsExist(err)
	var file *os.File
	if flag {
		//打开文件，
		fmt.Println("111")
		file, _ = os.OpenFile(logFileName, os.O_WRONLY|os.O_APPEND, 0666)
	} else {
		fmt.Println("222")
		//新建文件
		file, err = os.Create(logFileName)
		if err != nil {
			fmt.Println(err)
			return nil
		}
	}
	jl := &JsonLog{}
	jl.file = file
	//defer file.Close()
	enc := json.NewEncoder(file)
	jl.logger = enc
	return jl
}

func (jl *JsonLog) Log(data interface{}) {
	jl.logger.Encode(data)
}

func (jl *JsonLog) Close() {
	jl.file.Close()
}

//func CreateOrOpenFile(path string) (*os.File, error){
//
//	if _, flag := IsExists(path) ; flag && IsFile(path) {
//		return os.OpenFile(path, os.O_APPEND, 0666)
//	}
//}

// 写入文件,文件不存在则创建,如在则追加内容
func WriteFile(path string, str string) {
	_, b := IsFile(path)
	var f *os.File
	var err error
	if b {
		//打开文件，
		f, _ = os.OpenFile(path, os.O_APPEND, 0666)
	} else {
		//新建文件
		f, err = os.Create(path)
	}

	//使用完毕，需要关闭文件
	defer func() {
		err = f.Close()
		if err != nil {
			fmt.Println("err = ", err)
		}
	}()

	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	_, err = f.WriteString(str)
	if err != nil {
		fmt.Println("err = ", err)
	}
}

// 判断路径是否存在
func IsExists(path string) (os.FileInfo, bool) {
	f, err := os.Stat(path)
	return f, err == nil || os.IsExist(err)
}

// 判断所给路径是否为文件夹
func IsDir(path string) (os.FileInfo, bool) {
	f, flag := IsExists(path)
	return f, flag && f.IsDir()
}

// 判断所给路径是否为文件
func IsFile(path string) (os.FileInfo, bool) {
	f, flag := IsExists(path)
	return f, flag && !f.IsDir()
}
