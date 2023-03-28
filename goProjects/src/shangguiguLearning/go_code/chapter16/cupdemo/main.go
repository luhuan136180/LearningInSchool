package main

import (
	"fmt"
	"runtime"
)

func main() {
	cupNum:=runtime.NumCPU()
	fmt.Printf("CUPnMUM=",cupNum)
	//可以自己设置使用多个cpu
	//可以自己设置使用多个cpu
	//runtime.GOMAXPROCS(cpuNum - 1)
	//fmt.Println("ok")

	
}


