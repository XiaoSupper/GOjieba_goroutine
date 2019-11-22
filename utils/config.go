package utils

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
)

type Config struct {
	IsPos        bool   `json:"isPos"`
	GoroutineNum int    `json:"goroutineNum"`
	InputPath    string `json:"inputPath"`
	OutputPath   string `json:"outputPath"`
}

type JsonStruct struct {
}

var Configdata Config

var GOjieba = flag.String("GOjieba", "config.json", "gojieba配置文件")

func init() {
	flag.Parse()
	fmt.Println("-GOjieba:", *GOjieba)
	JsonParse := NewJsonStruct()
	//下面使用的是相对路径，config.json文件和main.go文件处于同一目录下
	JsonParse.Load(*GOjieba, &Configdata)

}

func NewJsonStruct() *JsonStruct {
	return &JsonStruct{}
}

func (jst *JsonStruct) Load(filename string, v interface{}) {
	//ReadFile函数会读取文件的全部内容，并将结果以[]byte类型返回
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, v)
	if err != nil {
		return
	}
}
