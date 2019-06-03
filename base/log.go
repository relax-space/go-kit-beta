package main

import (
	"log"
	"os"
)

//2019/06/03 15:39:35 log.go:8: hello log
func setLogTime() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

}

func setLogFile(fileName string) (err error) {
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return
	}
	log.SetOutput(f)
	return
}
