package main

import (
	"fmt"
)

func main() {
	c := ProjectDto{}

	err := ReadViper("", &c)
	fmt.Printf("read from file:%+v,err:%v", c, err)
	fmt.Println("\n============")
	p, err := ReadToDto()
	fmt.Printf("read from variable:%+v,err:%v", p, err)
}
