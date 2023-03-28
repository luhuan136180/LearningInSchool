package main

import (
	"fmt"
	"strings"
)

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Execption:", err)
		}
	}()
	var name string
	fmt.Scanln(&name)
	check(name)
}
func check(name string) {
	if strings.ContainsAny(name, ",#!") {
		panic("用户名含有非法字符")
	} else {
		fmt.Println(name)
	}
}
