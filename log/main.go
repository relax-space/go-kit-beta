package main

import (
	"github.com/astaxie/beego/logs"
)

/*
When the data reaches 1 million, a new file is automatically created. 当数据达到100万条之后，会自动创建新的文件
eg.  test.2019-07-29.001.log
*/
func main() {
	log := logs.NewLogger()
	log.SetLogger("file", `{"filename":"./logs/test.log"}`)
	for index := 0; index < 10000000; index++ {
		log.Error("111")
	}
}
