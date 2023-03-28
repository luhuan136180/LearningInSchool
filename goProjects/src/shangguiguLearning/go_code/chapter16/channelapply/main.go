package main

import (
	"fmt"
	"time"
)

func writeData(intChan chan int) {
	for i:=0;i<=50;i++{
		//放入数据
		intChan<-i*(i+1)*(i-5)
		fmt.Println("writeDatas=",i)
	}
	close(intChan)
}
func readData(intChan chan int, exitChan chan bool) {
	for{
		v,ok:=<-intChan
		if !ok{
			break
		}
		time.Sleep(time.Second)
		fmt.Printf("readData 读到数据=%v\n", v)
	}
	exitChan<-true
	close(exitChan)
}
func main() {
	intChan:=make(chan int,10)
	exitChan:=make(chan bool,1)

	go writeData(intChan)
	go readData(intChan,exitChan)
	//
	for{
		_,ok:=<-exitChan
		if !ok{
			break
		}
	}

}
