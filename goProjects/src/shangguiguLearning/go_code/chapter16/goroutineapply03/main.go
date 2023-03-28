package main

import (
	"fmt"
	"time"
)

func putNum(intChan chan int) {
	for i:=0;i<=80000;i++{
		intChan<-i
	}
	//
	close(intChan)
}
func primeNUm(primeChan chan int,intChan chan int,exitChan chan bool) {
	var flag bool
	for {
		num,ok:=<-intChan
		if !ok{
			break
		}
		flag=true//假设为素数
		for i:=2;i<num;i++{
			if num%i ==0{
				flag = false
				break
			}

		}
		if flag{
			primeChan<-num
		}
	}

	fmt.Println("有一个primeChan携程因为娶不到，退出")
	exitChan<-true
}
func main() {
	intChan :=make(chan int,2000)
	primeChan:=make(chan int,20000)
	//标识退出的管道
	exitChan:=make(chan bool,4)

	start := time.Now().Unix()
	go putNum(intChan)
	for i:=1;i<=4;i++{
		go primeNUm(primeChan,intChan,exitChan)

	}

	//
	for i:=0;i<4;i++{
		<-exitChan
	}
	end:=time.Now().Unix()
	fmt.Println("使用携程时=",end-start)
	close(primeChan)

	for{
		res,ok:=<-primeChan
		if !ok{
			break
		}
		fmt.Printf("素数=%d\n", res)
	}
	fmt.Println("main线程退出")
}
