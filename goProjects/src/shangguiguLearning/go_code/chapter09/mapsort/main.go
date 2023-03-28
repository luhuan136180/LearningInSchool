package main

import (
	"fmt"
	"sort"
)
//定义一个学生结构体
type Stu struct {
	Name string
	Age int
	Address string
}
func main() {
	map1 := make(map[int]int,10)
	map1[10]=12
	map1[1]=1
	map1[2]=20
	map1[3]=75
	map1[4]=12
	fmt.Println(map1)

	//如果按照map的key的顺序进行排序输出
	//1. 先将map的key 放入到 切片中
	//2. 对切片排序
	//3. 遍历切片，然后按照key来输出map的值
	var keys []int
	for k,_ := range map1{
		keys = append(keys,k)
	}
	sort.Ints(keys)
	fmt.Println(keys)
	//排序
	for _,v :=range keys{
	fmt.Printf("map[%v]=%v",v,map1[v])
	}

	//map的value 也经常使用struct 类型，
	//更适合管理复杂的数据(比前面value是一个map更好)，
	//比如value为 Student结构体 【案例演示，因为还没有学结构体，体验一下即可】
	//1.map 的 key 为 学生的学号，是唯一的
	//2.map 的 value为结构体，包含学生的 名字，年龄, 地址
	student :=make(map[string]Stu,10)
	sut1 :=Stu{"jack",12,"beijing"}
	stu2:=Stu{"tom",14,"jajaja"}
	student["no1"]=sut1
	student["no2"]=stu2
	fmt.Println(student)

	//
	for i,v := range student{
		fmt.Printf("Bianhaoshi=%v \t",i)
		fmt.Printf("name=%v",v.Name)
		fmt.Printf("age=%v",v.Age)
		fmt.Printf("aadress=%v",v.Address)
		fmt.Println()
	}

}
