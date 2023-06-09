package main

import "fmt"
//如果结构体的字段类型是: 指针，slice，和map的零值都是 nil ，即还没有分配空间
//如果需要使用这样的字段，需要先make，才能使用.
type Person struct {
	Name string
	Age string
	Scores [5]float64
	ptr *int
	slice []int
	map1 map[string]string
}
type monster struct {
	Name string
	Age int
}
func main() {
	var p1 Person
	fmt.Println(p1)
	if p1.ptr == nil{
		fmt.Println("ok1")
	}
	if p1.slice == nil{
		fmt.Println("ok2")
	}
	if p1.map1 == nil{
		fmt.Println("ok3")
	}
	//使用slice, 再次说明，一定要make
	p1.slice = make([]int,10)
	p1.slice[0]=100
	p1.slice[2]=123
	p1.slice[6]=122

	//使用map, 一定要先make
	p1.map1 = make(map[string]string,2)
	p1.map1["key1"]="tom~"
	fmt.Println(p1)

	//不同结构体变量的字段是独立，互不影响，一个结构体变量字段的更改，
	//不影响另外一个, 结构体是值类型
	var monster1 monster
	monster1.Name="niumo"
	monster1.Age=2000

	monster2:=monster1//结构体是值类型，默认为值拷贝
	monster2.Name="xiaomo"

	var monster3 monster
	monster3.Name="xiaoxiaohong"
	fmt.Println("monster1=", monster1) //monster1= {牛魔王 500}
	fmt.Println("monster2=", monster2) //monster2= {青牛精 500}
	fmt.Println("monster3",monster3)
}