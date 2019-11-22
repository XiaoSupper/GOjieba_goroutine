package main

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"runtime"
	"strings"
	"sync"
	"time"

	"./utils"
	"github.com/yanyiwu/gojieba"
)

var (
	DICT_DIR        string
	DICT_PATH       string
	HMM_PATH        string
	USER_DICT_PATH  string
	IDF_PATH        string
	STOP_WORDS_PATH string
)
var Seg *gojieba.Jieba

func init() {
	DICT_DIR = path.Join(path.Dir(getCurrentFilePath()), "dict")
	DICT_PATH = path.Join(DICT_DIR, "jieba.dict.utf8")
	HMM_PATH = path.Join(DICT_DIR, "hmm_model.utf8")
	USER_DICT_PATH = path.Join(DICT_DIR, "user.dict.utf8")
	IDF_PATH = path.Join(DICT_DIR, "idf.utf8")
	STOP_WORDS_PATH = path.Join(DICT_DIR, "stop_words.utf8")
}
func main() {

	var perNum int
	var ssFileList [][]string
	wg := sync.WaitGroup{}

	loadt := time.Now()
	Seg = gojieba.NewJieba(DICT_PATH, HMM_PATH, USER_DICT_PATH, IDF_PATH, STOP_WORDS_PATH)
	loadts := time.Since(loadt)

	fmt.Println("Loading...  Time-consuming:", loadts)
	defer Seg.Free()

	filelist := utils.GetFiles(utils.Configdata.InputPath)
	goNum := utils.Configdata.GoroutineNum
	if len(filelist) < goNum {
		goNum = len(filelist)
	}
	fmt.Println("GoRoutineNum:", goNum, "\nfileNum:", len(filelist))

	perNum = len(filelist) / goNum
	if len(filelist)%goNum != 0 && perNum == 0 {
		perNum = 1
	}

	for i := 0; i < goNum; i++ {
		ssFileList = append(ssFileList, filelist[perNum*i:perNum*(i+1)])
	}

	if len(filelist)%goNum != 0 {
		ssFileList = append(ssFileList, filelist[perNum*goNum:])
	}

	wg.Add(len(ssFileList))
	t := time.Now()
	for _, filelist := range ssFileList {
		go run(filelist, &wg)
	}

	wg.Wait()
	rept := time.Since(t)
	fmt.Println("ALL Done     Time-consuming:", rept)

}

func getCurrentFilePath() string {

	_, filePath, _, _ := runtime.Caller(1)
	return filePath

}
func run(filelist []string, wg *sync.WaitGroup) {

	if utils.Configdata.IsPos {
		for _, file := range filelist {
			runPosCut(file)
		}

	} else {
		for _, file := range filelist {
			runCut(file)
		}
	}
	wg.Done()
	return
}
func runPosCut(file string) {

	reg := regexp.MustCompile(".*<")
	t := time.Now()
	data := utils.Ioutil(path.Join(utils.Configdata.InputPath, file))

	fout, err := os.OpenFile(path.Join(utils.Configdata.OutputPath, file), os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Failed to open open file", err.Error())
		os.Exit(2)
	}
	defer fout.Close()

	for _, line := range strings.Split(data, "\n") {
		if len(strings.Replace(line, "\n", "", -1)) == 0 {
			continue
		}
		if reg.MatchString(line) {
			utils.WriteWithIo(fout, line+"\n")
		} else {
			words := Seg.Tag(line)
			// words := Seg.Tag(strings.Replace(line, "\n", "", -1))
			utils.WriteWithIo(fout, strings.TrimRight(strings.Join(words, " "), "/x"))
		}
	}

	rept := time.Since(t)
	fmt.Println(file, "Done     Time-consuming:", rept)
	return
}

func runCut(file string) {

	reg := regexp.MustCompile(".*<")
	t := time.Now()
	data := utils.Ioutil(path.Join(utils.Configdata.InputPath, file))

	fout, err := os.OpenFile(path.Join(utils.Configdata.OutputPath, file), os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Failed to open open file", err.Error())
		os.Exit(2)
	}
	defer fout.Close()
	for _, line := range strings.Split(data, "\n") {

		if reg.MatchString(line) {
			utils.WriteWithIo(fout, line+"\n")
		} else {
			words := Seg.Cut(line, true) //精确模式,第二个参数为true启用HMM新词发现
			// words := Seg.CutAll(line)    //全模式
			// words := Seg.CutForSearch(line, true) //搜索引擎模式
			utils.WriteWithIo(fout, strings.Join(words, " ")+"\n")
		}
	}

	rept := time.Since(t)
	fmt.Println(file, "Done     Time-consuming:", rept)
	return
}
