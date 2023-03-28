package main

import "fmt"

type Boy struct {
	no   int
	next *Boy
}

func Add(num int) (head *Boy) {
	temp := &Boy{}
	if num < 1 {
		return head
	}
	for i := 1; i <= num; i++ {
		boy := &Boy{
			no: i,
		}
		if i == 1 {
			head = boy
			temp = boy
			head.next = head
		} else {
			temp.next = boy
			temp = temp.next
			temp.next = head
		}
	}
	return head
}

func show(head *Boy) {
	if head == nil {
		fmt.Println("该链表为空")
		return
	}
	temp := head
	for true {
		fmt.Printf("小孩编号=%d ->", temp.no)
		if temp.next == head {
			break
		}
		temp = temp.next
	}
}
func test(startno int, head *Boy) bool {
	num := 0
	temp := head
	for true {
		if temp.next == head {
			num++
			break
		}
		temp = temp.next
		num++
	}
	if num > startno {
		return true
	} else {
		return false
	}
}
func PlayGame(head *Boy, StartNo int, num int) {
	//1. 空的链表我们单独的处理
	if head.next == nil {
		fmt.Println("空的链表，没有小孩")
		return
	}
	//留一个，判断 startNO <= 小孩的总数
	if !test(StartNo, head) {
		fmt.Println("StartNo值有误")
		return
	}

	temp := head
	helper := head
	for {
		if helper.next == head {
			break
		}
		helper = helper.next
	}
	for { //找到起始位置
		if temp.no == StartNo {
			break
		}
		temp = temp.next
		helper = helper.next
	}
	fmt.Println(temp.no)
	for {
		for i := 1; i < num; i++ {
			temp = temp.next
			helper = helper.next
		}
		//fmt.Println(temp.no)
		fmt.Printf("小孩编号为%d 出圈 \n", temp.no)
		helper.next = temp.next
		temp = temp.next
		if temp == helper {
			break
		}
	}
	fmt.Printf("小孩编号为%d 出圈 \n", temp.no)
}
func main() {
	first := Add(20)
	//显示
	show(first)
	PlayGame(first, 3, 3)
}
