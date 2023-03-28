package main

import (
	"fmt"
	"shangguiguLearning/go_code/chapter11/encapsualte/model"
)

func main() {
	person01 := model.NewPerson("smith")
	fmt.Println(person01)
	person01.SetAge(10)
	person01.SetSal(4000.0)
	fmt.Println(person01)
	fmt.Println(person01.Name,"age=",person01.GetAge(),"sal=",person01.GetSal())
	person02:= model.NewPerson02("tom",23,5000.0)
	fmt.Println(person02)
}