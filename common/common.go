package common

import (
	"fmt"
	"os"
	"reflect"
)

//写字符串到文件中
func WriteString(str string, fileName string) {
	dstFile,err := os.Create(fileName)
	if err != nil {
		fmt.Println(err); return
	}
	defer dstFile.Close()
	//s:="hello world"
	dstFile.WriteString(str+"\n")
}

//写入byte数据
func WriteByte(data []byte, fileName string) {
	dstFile,err := os.Create(fileName)
	if err != nil {
		fmt.Println(err); return
	}
	defer dstFile.Close()
	dstFile.Write(data)
}

func IsNil(i interface{}) bool {
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}

//创建文件夹目录
func Mkdir(_path string) (path string){

	res, err := PathExists(_path)
	if err != nil{
		panic("文件目录有问题请查验")
	}
	if res != true{
		_ = os.MkdirAll(_path, os.ModePerm)
	}
	return _path
}

func PathExists(path string) (bool, error) {
	/*
	   判断文件或文件夹是否存在
	   如果返回的错误为nil,说明文件或文件夹存在
	   如果返回的错误类型使用os.IsNotExist()判断为true,说明文件或文件夹不存在
	   如果返回的错误为其它类型,则不确定是否在存在
	*/
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
