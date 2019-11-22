package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func Ioutil(name string) string {

	if contents, err := ioutil.ReadFile(name); err == nil {
		//因为contents是[]byte类型，直接转换成string类型后会多一行空格,需要使用strings.Replace替换换行符
		return strings.Replace(string(contents), "\n", "", 1)
	}
	return "read file err!"
}

func WriteWithIo(fileObj *os.File, content string) {

	if _, err := io.WriteString(fileObj, content); err != nil {
		fmt.Println(err)
	}
}

func GetFiles(path string) []string {

	var filelist []string
	files, _ := ioutil.ReadDir(path)
	for _, file := range files {
		if file.IsDir() {
			continue
		} else {
			filelist = append(filelist, file.Name())
		}
	}
	return filelist
}
