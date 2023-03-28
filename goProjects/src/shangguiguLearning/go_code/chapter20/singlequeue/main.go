package main

import (
	"errors"
	"fmt"
	"os"
)

type Queue struct {
	maxSize int
	array [5]int
	front int
	rear int
}

func (this *Queue) addQueue(val int) (err error){
	if this.rear==this.maxSize-1{
		return errors.New("queue full")
	}
	this.rear++
	this.array[this.rear] = val
	return
}

func (this *Queue) GetQueue()(val int ,err error)  {
	if this.rear==this.front{
		//对空
		return -1,errors.New("queue empty")
	}
	this.front++
	val =this.array[this.front]
	return val,err
}

func (this *Queue) ShowQueue() {
	fmt.Println("队列当前的情况是:")
	for i:=this.front+1;i<=this.rear;i++{
		fmt.Printf("array[%d]=%d\t",i,this.array[i])
	}
	println()
}
func main() {
	queue:=Queue{
		maxSize: 5,
		front: -1,
		rear:-1,
	}
	var key string
	var val int
	for{
		fmt.Println("1. 输入add 表示添加数据到队列")
		fmt.Println("2. 输入get 表示从队列获取数据")
		fmt.Println("3. 输入show 表示显示队列")
		fmt.Println("4. 输入exit 表示显示队列")

		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入你要的入队列数")
			fmt.Scanln(&val)
			err:=queue.addQueue(val)
			if err!=nil{
				fmt.Println(err.Error())
			}else {
				fmt.Println("加入队列ok")
			}
		case "get":
			val,err:=queue.GetQueue()
			if err!=nil{
				fmt.Println(err.Error())
			}else{
				fmt.Println("从队列中取出了一个数=", val)
			}
		case "show":
			queue.ShowQueue()
		case "exit":
			os.Exit(0)
		}
	}
}