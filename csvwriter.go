// json.go
package gotools

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Str interface {
	String() string
}

type CSVWriter struct {
	file *os.File
}

func NewCSVWriter(logFileName string) (*CSVWriter, error) {
	_, err := os.Stat(logFileName)
	flag := err == nil || os.IsExist(err)
	var file *os.File
	if flag {
		//打开文件，
		fmt.Println("File Exists")
		file, _ = os.OpenFile(logFileName, os.O_WRONLY|os.O_APPEND, 0666)
	} else {
		fmt.Println("File Created")
		//新建文件
		file, err = os.Create(logFileName)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}
	jl := &CSVWriter{}
	jl.file = file
	return jl, nil
}

func (jl *CSVWriter) write(data Str) {
	jl.file.Write([]byte(data.String() + "\n"))
}

func (jl *CSVWriter) writeStr(data string) {
	jl.file.Write([]byte(data + "\n"))
}

func (jl *CSVWriter) Close() {
	jl.file.Close()
}

type CSVReader struct {
	file *os.File
}

func NewCSVReader(logFileName string) (*CSVReader, error) {
	_, err := os.Stat(logFileName)
	flag := err == nil || os.IsExist(err)
	var file *os.File
	if flag {
		//打开文件，
		fmt.Println("File Exists")
		file, _ = os.OpenFile(logFileName, os.O_RDONLY, 0666)
	} else {
		fmt.Println("File Created")
		//新建文件
		file, err = os.Create(logFileName)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}
	jl := &CSVReader{}
	jl.file = file
	return jl, nil
}

func (jl *CSVReader) Read() []int {
	reader := bufio.NewReader(jl.file)
	ret := make([]int, 0)
	for {
		buf, err := reader.ReadString('\n')
		buf = strings.Trim(buf, "\n")
		fmt.Println(buf)
		if err == io.EOF {
			break
		}
		i, err := strconv.Atoi(buf)
		if err != nil {
			fmt.Println(err)
		}
		ret = append(ret, i)
	}
	return ret
}

func (jl *CSVReader) Close() {
	jl.file.Close()
}
