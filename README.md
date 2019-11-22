# GOjieba_goroutine
多协程调用GOjieba分词

## 简介

+ 支持多协程调用GOjieba分词
+ 核心算法底层由C++实现，GO封装，性能高效。
+ 字典路径可配置，NewJieba(...string)可变形参，当参数为空时使用默认词典(默认路径在GOjieba包内，使用默认词典需要先安装GOjiaba(go get github.com/yanyiwu/gojieba))

## 用法

```
go get github.com/yanyiwu/gojieba
```
run.bat
```
选项 1 编译运行main.go
选项 2 编译main.go
选项 3 运行main.exe
```

## Benchmark
### 测试环境

```
操作系统：win10
内核：4
逻辑处理器：8
```
### 测试语料

```
8个20.2M语料文件
```
### 测试结果

```
-GOjieba: config.json
Loading...  Time-consuming: 2.1719415s
GoRoutineNum: 4
fileNum: 8

D0_tmp_out4 Done     Time-consuming: 7.586548s
D0_tmp_out6 Done     Time-consuming: 7.586548s
D0_tmp_out0 Done     Time-consuming: 7.6664455s
D0_tmp_out2 Done     Time-consuming: 7.7681726s
D0_tmp_out7 Done     Time-consuming: 9.6924108s
D0_tmp_out5 Done     Time-consuming: 9.8260509s
D0_tmp_out1 Done     Time-consuming: 10.0413262s
D0_tmp_out3 Done     Time-consuming: 10.0143987s
ALL Done     Time-consuming: 17.7845665s
```

## Contact

+ Email: `blcuxiao@126.com`

[CppJieba]:http://github.com/yanyiwu/cppjieba
[cppjiebaapp多线程调用CppJieba]:https://github.com/xiabo0816/cppjiebaapp
[GoJieba]:http://github.com/yanyiwu/gojieba
[Jieba]:https://github.com/fxsjy/jieba
[Jieba中文分词系列性能评测]:http://yanyiwu.com/work/2015/06/14/jieba-series-performance-test.html
