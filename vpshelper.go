package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	b, err := ioutil.ReadFile("/etc/v2ray/config.json")
	if err != nil {
		fmt.Println("Error occur when read file")
	}
	//fmt.Println(b)
	str := string(b)
	fmt.Println(str)
}
