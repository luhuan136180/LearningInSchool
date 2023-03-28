package main

import "fmt"

func main() {
	//演示map切片的使用
	/*
		要求：使用一个map来记录monster的信息 name 和 age, 也就是说一个
		monster对应一个map,并且妖怪的个数可以动态的增加=>map切片
	*/
	//1. 声明一个map切片————即一个类型为map的切片
	var a map[string]string
	a=make(map[string]string,2)
	fmt.Println(a)
	var arr []int=make([]int,2,10)
	arr = []int{1,23,4}
	fmt.Println(arr)
	var monster []map[string]string

	monster =make([]map[string]string,2)
	if monster[0] ==nil{
		monster[0]=make(map[string]string,2)//第一种map创建方式
		monster[0]["name"]="niumowang"
		monster[0]["age"]="600"
	}

	if monster[1] ==nil{
		monster[1]=map[string]string{
			"name":"yutu",
			"age":"500",
		}
	}

	fmt.Println(monster)

	//当需要动态添加的时候
	//这里我们需要使用到切片的append函数，可以动态的增加monster
	//1. 先定义个monster信息
	newMonster :=map[string]string{//使用第三种map创建方式
		"name":"huli",
		"age":"200",
	}
	//第二种map创建方式
	newMonster2 :=make(map[string]string,2)
	newMonster2["name"]="zhu"
	newMonster2["age"]="100"

	monster =append(monster,newMonster)
	fmt.Println(newMonster)
	fmt.Println(monster)

	monster=append(monster,newMonster2)
	fmt.Println(monster)
}
