package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//1)创建一个byte类型的26个元素的数组，分别 放置'A'-'Z‘。
	//使用for循环访问所有元素并打印出来。提示：字符数据运算 'A'+1 -> 'B'
	//思路
	//1. 声明一个数组 var myChars [26]byte
	//2. 使用for循环，利用 字符可以进行运算的特点来赋值 'A'+1 -> 'B'
	//3. 使用for打印即可
	//代码:
	var myChars [26]byte//存储的时候，值用的整数
	for i:=0;i<len(myChars);i++{
		myChars[i]='A'+byte(i)// 注意需要将 i => byte;i本身是int类型
	}
	for i := 0; i < 26; i++ {
		fmt.Printf("%c ", myChars[i])//输出时需要将整数值转换成对应的码值

	}
	fmt.Println("")
	//请求出一个数组的最大值，并得到对应的下标

	//思路
	//1. 声明一个数组 var intArr[5] = [...]int {1, -1, 9, 90, 11}
	//2. 假定第一个元素就是最大值，下标就0
	//3. 然后从第二个元素开始循环比较，如果发现有更大，则交换
	var intArr [5]int=[5]int{1,-1,9,90,11}
	maxindex := 0
	maxvalue :=intArr[maxindex]
	for i:=0;i<len(intArr);i++{
		if intArr[i]>=maxvalue {
			maxindex =i
			maxvalue=intArr[maxindex]
		}
	}
	fmt.Println("数组中最大值=%v",maxvalue)

	//请求出一个数组的和和平均值。for-range//运用不熟练
	//思路
	//1. 就是声明一个数组  var intArr[5] = [...]int {1, -1, 9, 90, 11}
	//2. 求出和sum
	//3. 求出平均值
	//代码
	var intArr02 [5]int= [5]int {1, -1, 9, 93, 11}
	sum:=0
	for _,value :=range intArr02{//for..range 遍历数组
		sum +=value
	}
	fmt.Printf("sum=%v,平均值=%v\n",sum,float64(sum)/float64(len(intArr02)))

	//要求：随机生成五个数，并将其反转打印
	//思路
	//1. 随机生成五个数 , rand.Intn() 函数
	//2. 当我们得到随机数后，就放到一个数组 int数组
	//3. 反转打印 , 交换的次数是  len / 2, 倒数第一个和第一个元素交换, 倒数第2个和第2个元素交换
	var intArr03 [6]int
	rand.Seed(time.Now().UnixNano())//time.Now().UnixNano():取时间戳
	for i:=0;i<len(intArr03);i++{
		intArr03[i]=rand.Intn(100)
	}
	var temp int=0
	fmt.Println(intArr03)
	for i:=0;i<len(intArr03)/2;i++{
		temp = intArr03[i]
		intArr03[i]=intArr03[len(intArr03)-1-i]
		intArr03[len(intArr03)-1-i] = temp
	}
	fmt.Println("\"交换后~=",intArr03)
}
