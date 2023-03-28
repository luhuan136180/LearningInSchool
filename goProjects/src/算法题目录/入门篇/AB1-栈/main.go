package main

import (
	"errors"
	"fmt"
)

type Stack []interface{}

func (stack *Stack) Len() int {
	return len(*stack)
}

func (stack *Stack) Push(x interface{}) {
	*stack = append(*stack, x)
}
func (stack *Stack) Pop() (interface{}, error) {
	n := stack.Len()
	thisStack := *stack //生成临时栈
	if n == 0 {
		return nil, errors.New("error")
	}
	u := thisStack[n-1]
	*stack = thisStack[:n-1]
	return u, nil
}
func (stack *Stack) Top() (interface{}, error) {
	n := stack.Len()
	thisStack := *stack
	if n == 0 {
		return nil, errors.New("error")
	}
	return thisStack[n-1], nil
}

func main() {
	var n int
	fmt.Scanln(&n)
	var stack Stack
	for i := 0; i < n; i++ {
		var str string
		fmt.Scan(&str)
		switch str {
		case "push":
			var v interface{}
			var m int
			fmt.Scan(&m)
			v = m
			stack.Push(v)
		case "top":
			u, err := stack.Top()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(u)
			}
		case "pop":
			u, err := stack.Pop()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(u)
			}
		}
	}
}
