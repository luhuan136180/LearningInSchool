package main

import "fmt"

type HeroNode struct {
	no       int
	name     string
	nickname string
	pre      *HeroNode //这个表示指向前一个结点
	next     *HeroNode //这个表示指向下一个结点
}

//给双向链表插入一个结点
//编写第一种插入方法，在单链表的最后加入.[简单]
func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	temp := head
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	//3. 将newHeroNode加入到链表的最后
	temp.next = newHeroNode
	newHeroNode.pre = temp
}

//给链表插入一个结点
//编写第2种插入方法，根据no 的编号从小到大插入..【实用】
func InsertHeroNode2(head *HeroNode, newheroname *HeroNode) {
	temp := head
	flag := true
	for {
		if temp.next == nil {
			break
		} else if temp.next.no > newheroname.no {
			break
		} else if temp.next.no == newheroname.no {
			flag = false
			break
		}
		temp = temp.next
	}
	if flag {
		newheroname.next = temp.next
		newheroname.pre = temp
		if temp.next != nil {
			temp.next.pre = newheroname
		}
		temp.next = newheroname

	} else {
		fmt.Println("对不起，已经存在no=", newheroname.no)
		return
	}
}

//删除一个结点
func DelHerNode(head *HeroNode, id int) {
	temp := head
	flag := true
	for {
		if temp.next == nil {
			fmt.Println("没有找到该节点")
			break
		} else if temp.next.no == id {
			flag = false
			break
		}
		temp = temp.next
	}
	if flag {
	} else {
		temp.next = temp.next.next
		if temp.next != nil {
			temp.next.pre = temp
		}
	}
}

//显示链表的所有结点信息
//这里仍然使用单向的链表显示方式
func ListHeroNode(head *HeroNode) {

	//1. 创建一个辅助结点[跑龙套, 帮忙]
	temp := head

	// 先判断该链表是不是一个空的链表
	if temp.next == nil {
		fmt.Println("空空如也。。。。")
		return
	}
	//2. 遍历这个链表
	for {
		fmt.Printf("[%d , %s , %s]==>", temp.next.no,
			temp.next.name, temp.next.nickname)
		//判断是否链表后
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
}

//逆序打印
func ListHeroNode2(head *HeroNode) {

	//1. 创建一个辅助结点[跑龙套, 帮忙]
	temp := head

	// 先判断该链表是不是一个空的链表
	if temp.next == nil {
		fmt.Println("空空如也。。。。")
		return
	}

	//2. 让temp定位到双向链表的最后结点
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}

	//2. 遍历这个链表
	for {
		fmt.Printf("[%d , %s , %s]==>", temp.no,
			temp.name, temp.nickname)
		//判断是否链表头
		temp = temp.pre
		if temp.pre == nil {
			break
		}
	}
}

func main() {
	//1. 先创建一个头结点,
	head := &HeroNode{}

	//2. 创建一个新的HeroNode
	hero1 := &HeroNode{
		no:       1,
		name:     "宋江",
		nickname: "及时雨",
	}

	hero2 := &HeroNode{
		no:       2,
		name:     "卢俊义",
		nickname: "玉麒麟",
	}

	hero3 := &HeroNode{
		no:       3,
		name:     "林冲",
		nickname: "豹子头",
	}
	InsertHeroNode2(head, hero3)
	InsertHeroNode2(head, hero1)
	fmt.Println()
	ListHeroNode(head)
	InsertHeroNode2(head, hero2)
	fmt.Println()
	ListHeroNode(head)
	fmt.Println()
	ListHeroNode2(head)
	DelHerNode(head, 2)
	fmt.Println()
	ListHeroNode(head)
}
