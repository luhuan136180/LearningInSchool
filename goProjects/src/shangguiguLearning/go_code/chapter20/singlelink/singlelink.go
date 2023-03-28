package main

import "fmt"

type HeroNode struct {
	no       int
	name     string
	nickname string
	next     *HeroNode
}

//给链表插入一个结点
//编写第一种插入方法，在单链表的最后加入.[简单]
func InsertHeroNode(head *HeroNode, newheroname *HeroNode) {
	//思路
	//1. 先找到该链表的最后这个结点
	//2. 创建一个辅助结点[跑龙套, 帮忙]temp
	temp := head
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next // 让temp不断的指向下一个结点
	}
	//3. 将newHeroNode加入到链表的最后
	temp.next = newheroname
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
	}
}

//显示链表的所有结点信息
func ListHeroNode(head *HeroNode) {
	temp := head
	if temp.next == nil {
		fmt.Println("该列表为空")
		return
	}
	for {
		fmt.Printf("[%d,%s,%s]=>", temp.next.no, temp.next.name, temp.next.nickname)
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
}
func main() {
	//1.创建头结点
	head := &HeroNode{}
	//2.创建一个heroname
	hero1 := &HeroNode{
		name:     "songjian",
		no:       1,
		nickname: "及时雨",
	}
	//fmt.Println(&hero1)	0xc000006028
	//fmt.Println(hero1)	&{1 songjian 及时雨 <nil>}
	//fmt.Println(*hero1)	{1 songjian 及时雨 <nil>}
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
	InsertHeroNode2(head, hero2)
	ListHeroNode(head)
	// 5 删除
	fmt.Println()
	DelHerNode(head, 2)
	ListHeroNode(head)
	fmt.Println()
	DelHerNode(head, 3)
	ListHeroNode(head)
}
