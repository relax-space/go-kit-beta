package main

import (
	"fmt"
	"log"
)

func main() {

	fmt.Println(Unique([]string{"1", "2", "1", "4", "1", "8", "1"}))

	setLogTime()
	log.Println("hello log")

	if err := setLogFile("test.log"); err != nil {
		fmt.Println(err)
	}
	log.Println("111")
}
