package main

import "fmt"

type CatNode struct {
	no   int
	name string
	next *CatNode
}

func InsertCatNode(head *CatNode, newCatNode *CatNode) {
	if head.next == nil {
		head.no = newCatNode.no
		head.name = newCatNode.name
		head.next = head
		return
	}
	temp := head
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}
	temp.next = newCatNode
	newCatNode.next = head
}

//删除一只猫
func DelCatNode(head *CatNode, id int) *CatNode {

	temp := head
	helper := head
	//空链表
	if temp.next == nil {
		fmt.Println("这是一个空的环形链表，不能删除")
		return head
	}

	//如果只有一个结点
	if temp.next == head { //只有一个结点
		if temp.no == id {
			temp.next = nil
		}
		return head
	}

	//将helper 定位到链表最后
	for {
		if helper.next == head {
			break
		}
		helper = helper.next
	}

	//如果有两个包含两个以上结点
	flag := true
	for {
		if temp.next == head { //如果到这来，说明我比较到最后一个【最后一个还没比较】
			break
		}
		if temp.no == id {
			if temp == head { //说明删除的是头结点
				head = head.next
			}
			//恭喜找到., 我们也可以在直接删除
			helper.next = temp.next
			fmt.Printf("猫猫=%d\n", id)
			flag = false
			break
		}
		temp = temp.next     //移动 【比较】
		helper = helper.next //移动 【一旦找到要删除的结点 helper】
	}
	//这里还有比较一次
	if flag { //如果flag 为真，则我们上面没有删除
		if temp.no == id {
			helper.next = temp.next
			fmt.Printf("猫猫=%d\n", id)
		} else {
			fmt.Printf("对不起，没有no=%d\n", id)
		}
	}
	return head
}

func DelCatNode2(head *CatNode, no int) (newhead *CatNode) {
	temp := head
	helper := head
	flag := false
	if head == nil {
		fmt.Println("该链表为空")
		return nil
	}
	for true {
		if helper.next == head {
			break
		}
		helper = helper.next
	}

	if head.no == no {
		head = head.next
		helper.next = head
		return head
	} else {
		for true {
			if temp.next == head {
				if temp.no == no {
					flag = true
				}
				break
			} else if temp.no == no {
				flag = true
				break
			}
			temp = temp.next
			helper = helper.next
		}
		if flag {
			helper.next = temp.next
			return head
		} else {
			return head
		}
	}
}

func ListCircleLink(head *CatNode) {
	temp := head
	if temp.next == nil {
		fmt.Println("该链表为空")
		return
	}
	for true {
		fmt.Printf("猫的信息为=[id=%d name=%s] ->", temp.no, temp.name)
		if temp.next == head {
			break
		}
		temp = temp.next
	}
}

func main() {
	//这里我们初始化一个环形链表的头结点
	head := &CatNode{} //var head *CatNode = &CatNode{}

	//创建一只猫
	cat1 := &CatNode{
		no:   1,
		name: "tom",
	}
	cat2 := &CatNode{
		no:   2,
		name: "tom2",
	}
	cat3 := &CatNode{
		no:   3,
		name: "tom3",
	}
	InsertCatNode(head, cat1)
	InsertCatNode(head, cat2)
	InsertCatNode(head, cat3)
	//fmt.Println(head)
	ListCircleLink(head)
	fmt.Println()
	head = DelCatNode2(head, 3)

	fmt.Println()
	ListCircleLink(head)

}
