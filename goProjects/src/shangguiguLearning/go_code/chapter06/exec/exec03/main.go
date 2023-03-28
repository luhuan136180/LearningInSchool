package main

import "fmt"
//题3：猴子吃桃子问题有一堆桃子，猴子第一天吃了其中的一半，
//并再多吃了一个！以后每天猴子都吃其中的一半，然后再多吃一个。
//当到第十天时，想再吃时（还没吃），发现只有1个桃子了。问题：最初共多少个桃子？

//思路分析
/*
1)第10天只有一个桃子
2)第9天有几个桃子  =  (第10天桃子数量 + 1) * 2
3)规律: 第n天的桃子数据  peach(n) = (peach(n+1) + 1) * 2
*/
func peach(n int)int{
	if n==10{
		return 1
	}else{
		return (peach(n+1)+1)*2
	}

}
func main() {
	fmt.Println("第1天的桃子个数",peach(1))
}