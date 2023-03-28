package main
//联系接口编程
import (
	"fmt"
	"math/rand"
	"sort"
)

type Hero struct {
	Name string
	Age int
}
type HeroSlice []Hero


func (hs HeroSlice) Len() int{
	return len(hs)
}
//Less方法就是决定你使用什么标准进行排序
//1. 按Hero的年龄从小到大排序!!
func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age<hs[j].Age
}
func (hs HeroSlice) Swap(i, j int) {
	hs[i],hs[j] = hs[j],hs[i]
}
//

type student struct{
	Name string
	Age int
	Score int
}
type StuSlice []student
func (stu StuSlice) Len() int {
	return len(stu)
}
func (stu StuSlice) Less(i, j int) bool {
	return stu[i].Score<stu[j].Score
}
func (stu StuSlice)Swap(i,j int)  {
	stu[i],stu[j]=stu[j],stu[i]
}
//声明一个hero结构体类型
func main() {
	var intSlice = []int{0,13,20,22,11,-1,-4,37}
	sort.Ints(intSlice)
	fmt.Println(intSlice)


	var heroes HeroSlice
	for i:=0;i<10;i++{
		hero:=Hero{
			Name: fmt.Sprintf("英雄%d",rand.Intn(100)),
			Age: rand.Intn(100),
		}
		heroes = append(heroes,hero)
	}
	for _,value :=range heroes{
		fmt.Println(value)
	}
	sort.Sort(heroes)
	fmt.Println("-----------排序后------------")
	for _,v :=range heroes{
		fmt.Println(v)
	}
	fmt.Println("-----------lianxi------------")
	var stuSlices StuSlice
	for i:=0;i<10;i++{
		stu:=student{
			Name: fmt.Sprintf("xuesheng%d",rand.Intn(100)),
			Age: rand.Intn(100),
			Score: rand.Intn(100),
		}
		stuSlices=append(stuSlices,stu)
	}
	for _,v :=range stuSlices{
		fmt.Println(v)
	}
	sort.Sort(stuSlices)
	fmt.Println("-----------排序后------------")
	for _,v :=range stuSlices{
		fmt.Println(v)
	}
}


