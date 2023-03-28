package main

import "fmt"

func main()  {
	for i:=1;i<=10;i++{
		fmt.Println("你好，尚硅谷！",i)
	}

	//for的第二种写法
	j:=1
	for j<+10{
		fmt.Println("你好，世界！",j)
		j++//迭代变量的迭代
	}

	//for循环的第三种写法, 这种写法通常会配合break使用
	//这里也等价 for ; ; {
	k :=1
	for {
		if k <=10{
			fmt.Println("OK2!~~",k)
		}else {
			break//终止循环，跳出
		}
		k++
	}

	//	//字符串遍历方式1-传统方式
	//	// var str string = "hello,world!北京"
	//	// for i := 0; i < len(str); i++ {
	//	// 	fmt.Printf("%c \n", str[i]) //使用到下标...
	//	// }

	var str string = "hello,world.北京"//中文一个汉字对应3个字节，数组按照单字节储存
	str2 := []rune(str)//吧str 转成 【】rune
	for i:=0;i<len(str2);i++{
		fmt.Printf("%c \n",str2[i])
	}

	fmt.Println()
	//字符串遍历方式2-for-range
	//按照字符方式遍历
	str = "abc1dfg 伤害"
	for index,val := range str{
		fmt.Printf("index=%d,val=%c \n",index,val)
	}
	}
