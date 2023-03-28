package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var count int
	rand.Seed(time.Now().Unix())
	// Seed should not be called concurrently with any other Rand method.
	for{
		n:= rand.Intn(100)+1//func Intn: func Intn(n int) int
		// 返回一个取值范围在[0,n)的伪随机int值，如果n<=0会panic。
		fmt.Println("n=",n)
		count++
		if(n==99){
			break
		}
	}
	fmt.Println("生成99一共用了，",count)

	////这里演示一下指定标签的形式来使用 break
	lable2:
		for i:=0;i<4;i++{
			for j:=0;j<10;j++{
				if j==2{
					break lable2
				}
				fmt.Println("j=",j)
			}
		}
}
