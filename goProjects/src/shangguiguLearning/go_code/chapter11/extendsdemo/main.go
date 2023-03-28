package main

import "fmt"
type Student struct {
	Name string
	Age int
	Score int
}
//将Pupil 和 Graduate 共有的方法也绑定到 *Student
func(student *Student) showInfo() {
	fmt.Printf("学生姓名=%v,年龄=%v 成绩=%v",student.Name,student.Age,student.Score)
}
func(stu *Student) SetScore(score int) {
	stu.Score=score
}
//给 *Student 增加一个方法，那么 Pupil 和 Graduate都可以使用该方法
func(stu *Student) GetSum(n1 int, n2 int) int{
	return n1+n2
}

type Pupil struct {
	Student//嵌入了Student匿名结构体
}
//这时Pupil结构体特有的方法，保留
func(p *Pupil) testing() {
	fmt.Println("xiaoxuesheng正在考试ing")
}

//daxuesheng
type Graduate struct {
		Student
}
//这时Graduate结构体特有的方法，保留
func(stu *Graduate) testing() {
	fmt.Println("daxuesheng正在考试")
}

func main() {
	//当我们对结构体嵌入了匿名结构体使用方法会发生变化
	pupil01:=&Pupil{}
	pupil01.Student.Name="tom"
	pupil01.Age=6
	pupil01.testing()
	pupil01.Student.SetScore(69)
	pupil01.showInfo()
	pupil01.GetSum(2,4)

	//
	stu02:=&Graduate{}
	stu02.Student.Name = "mary~"
	stu02.Student.Age = 28
	stu02.testing()
	stu02.Student.SetScore(90)
	stu02.Student.showInfo()
	fmt.Println("res=",stu02.Student.GetSum(10,30))
}