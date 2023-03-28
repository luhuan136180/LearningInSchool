package main

import "fmt"

//8.1班上有30个学⽣，
//每个学⽣的信息包括学号、姓名、性别、年龄、三门课的成绩。
//要求建⽴学⽣信息的结构体student，输⼊这30个
//学⽣的信息，然后打印输出各项数据。

type Student struct {
	num   string
	name  string
	sex   string
	age   int
	score [3]float64
}

func main() {
	var stu [30]Student
	for i := 0; i < len(stu); i++ {
		fmt.Scan(&stu[i].num, &stu[i].name, &stu[i].sex, &stu[i].age, &stu[i].score[0], &stu[i].score[1], &stu[i].score[2])
	}
	for _, v := range stu {
		fmt.Println(v.print)
	}
}
func (stu *Student) print() {
	fmt.Println(" 学 号 :", stu.num)
	fmt.Println(" 姓 名 :", stu.name)
	fmt.Println(" 年 龄 ：", stu.age)
	fmt.Println(" 性 别 ：", stu.sex)
	fmt.Println(" 数 学 :", stu.score[0])
	fmt.Println(" 语 ⽂ ：", stu.score[1])
	fmt.Println("英语：", stu.score[2], "\n")
}