package main

import (
	"fmt"
	"strconv"
	"time"
)

func test03(){
	str :=""
	for i:=0;i<100000;i++{
		str +="hello"+strconv.Itoa(i)
	}
}
func main() {
	start := time.Now().Unix()
	test03()
	end := time.Now().Unix()
	fmt.Printf("执行test03（）耗费%v秒",end-start)
}