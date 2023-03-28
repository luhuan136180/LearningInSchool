package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Max struct {
	temp int
	x int
	y int
}
func main() {
	//生成随机二维数组
	res:=[5][6]int{}
	rand.Seed(time.Now().UnixNano())
	for i:=0;i<len(res);i++ {
		for j:=0;j<len(res[i]);j++{
			res[i][j]=rand.Intn(100)
		}
	}
	fmt.Println(res)

	for i := 0; i < len(res); i++ {
		var max Max
		max.temp=0
		for j := 0; j < len(res[i]); j++ {//找每行最大
			if res[i][j] > max.temp{
				max.temp=res[i][j]
				max.x=i
				max.y=j
			}
		}
		//判断是否为所在排最大
		maxfinal:=findpaimax(max,res)
		if maxfinal.temp!=-1 {
			fmt.Printf("鞍数: %d ,行数：%d,排数：%d \n",maxfinal.temp,maxfinal.x,maxfinal.y)
		}
	}
}
func findpaimax(max Max,res [5][6]int) Max {
	j:=max.y
	temp:=max.temp
	for i:=0;i<len(res);i++ {
		if i==max.x {
			continue
			fmt.Println("aaaa")
		}
		if res[i][j]>temp {
			max.temp=-1
			break
		}
	}
	return max
}
