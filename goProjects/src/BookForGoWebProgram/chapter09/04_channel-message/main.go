package main

import (
	"fmt"
	"time"
)

func thrower(c chan int) {
	for i := 0; i < 6; i++ {
		c <- i
		fmt.Println("Threw >>", i)
	}
}

func catcher(c chan int) {
	for i := 0; i < 6; i++ {
		num := <-c
		fmt.Println("Caught <<", num)
	}
}

//无缓冲通道
//func main() {
//	c := make(chan int)
//	go thrower(c)
//	go catcher(c)
//	time.Sleep(10 * time.Second)
//}

//有缓冲通道
func main() {
	c := make(chan int, 2)
	go thrower(c)
	go catcher(c)
	time.Sleep(10 * time.Second)
}
