package main

import (
	"fmt"
	"os"
)

type Emp struct {
	Id   int
	Name string
	Next *Emp
}

func (this *Emp) showMe() {
	fmt.Printf("链表%d 找到该雇员 %d\n", this.Id%7, this.Id)

}

type EmpLink struct {
	Head *Emp
}

func (this *EmpLink) Insert(emp *Emp) {
	temp := this.Head
	helper := &Emp{}
	if this.Head == nil {
		this.Head = emp
		return
	}
	for {
		if temp != nil {
			if temp.Id > emp.Id {
				break
			}
			helper = temp
			temp = temp.Next
		} else {
			break
		}
	}
	emp.Next = temp
	helper.Next = emp
}

func (this *EmpLink) ShowLink(no int) {
	if this.Head == nil {
		fmt.Println("链表为空")
		return
	}

	//
	cur := this.Head
	for {
		if cur != nil {
			fmt.Printf("链表%d 雇员id=%d 名字=%s ->", no, cur.Id, cur.Name)
			cur = cur.Next
		} else {
			break
		}
	}
	fmt.Println()
}

func (this *EmpLink) FindById(id int) *Emp {
	cur := this.Head
	for {
		if cur != nil && cur.Id == id {
			return cur
		} else if cur == nil {
			break
		}
		cur = cur.Next
	}
	return nil
}

type HashTable struct {
	LinkArr [7]EmpLink
}

func (this *HashTable) Insert(emp *Emp) {
	linkNo := this.HashFun(emp.Id)
	this.LinkArr[linkNo].Insert(emp)
}
func (this *HashTable) HashFun(id int) int {
	return id % 7
}

func (this *HashTable) ShowAll() {
	for i := 0; i < len(this.LinkArr); i++ {
		this.LinkArr[i].ShowLink(i)
	}
}

func (this *HashTable) FindById(id int) *Emp {
	linkNo := this.HashFun(id)
	return this.LinkArr[linkNo].FindById(id)
}

func main() {
	key := ""
	id := 0
	name := ""
	var hashtable HashTable
	for true {
		fmt.Println("===============雇员系统菜单============")
		fmt.Println("input 表示添加雇员")
		fmt.Println("show  表示显示雇员")
		fmt.Println("find  表示查找雇员")
		fmt.Println("exit  表示退出系统")
		fmt.Println("请输入你的选择")
		fmt.Scanln(&key)
		switch key {
		case "input":
			fmt.Println("输入雇员id")
			fmt.Scanln(&id)
			fmt.Println("输入雇员name")
			fmt.Scanln(&name)
			emp := Emp{
				Id:   id,
				Name: name,
			}
			hashtable.Insert(&emp)
		case "show":
			hashtable.ShowAll()
		case "find":
			fmt.Println("请输入id号:")
			fmt.Scanln(&id)
			emp := hashtable.FindById(id)
			if emp == nil {
				fmt.Printf("id=%d 的雇员不存在\n", id)
			} else {
				emp.showMe()
			}
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("输入有误")
		}
	}
}
