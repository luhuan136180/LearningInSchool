package main

import (
	"errors"
	"fmt"
)

//使用数组来模拟一个栈的使用
type Stack struct {
	MaxTop int
	Top    int
	arr    [5]int
}

func (this *Stack) Push(val int) (err error) {
	//
	if this.Top == this.MaxTop-1 {
		fmt.Println("stack full")
		return errors.New("stack if full")
	}
	this.Top++
	this.arr[this.Top] = val
	return
}
func (this *Stack) Pop() (val int, err error) {
	//判断是否为空
	if this.Top == -1 {
		fmt.Println("stack empty")
		return 0, errors.New("stack empty")
	}
	val = this.arr[this.Top]
	this.Top--
	return val, nil
}
func (this *Stack) List() {
	//
	if this.Top == -1 {
		fmt.Println("stack empty")
		return
	}
	fmt.Println("栈的情况如下：")
	for i := this.Top; i >= 0; i-- {
		fmt.Printf("arr[%d]=%d\n", i, this.arr[i])
	}
}
func main() {
	stack := &Stack{
		MaxTop: 5,
		Top:    -1,
	}
	//入栈
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)

	//显示
	stack.List()
	val, _ := stack.Pop()
	fmt.Println("出栈val=", val) // 5
	//显示
	stack.List() //

	fmt.Println()
	val, _ = stack.Pop()
	val, _ = stack.Pop()
	val, _ = stack.Pop()
	val, _ = stack.Pop()
	val, _ = stack.Pop()       // 出错
	fmt.Println("出栈val=", val) // 5
	//显示
	stack.List() //
}
