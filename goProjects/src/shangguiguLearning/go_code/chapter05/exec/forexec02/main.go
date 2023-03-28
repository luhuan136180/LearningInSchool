package main
import "fmt"

func main() {

	//1)统计3个班成绩情况，每个班有5名同学，
	//求出各个班的平均分和所有班级的平均分[学生的成绩从键盘输入]

	//分析实现思路
	//1. 统计1个班成绩情况，每个班有5名同学, 求出该班的平均分【学生的成绩从键盘输入】=》先易后难
	//2. 学生数就是5个 [先死后活]
	//3. 声明一个sum 统计班级的总分

	//分析实现思路2
	//1. 统计3个班成绩情况，每个班有5名同学, 求出每个班的平均分【学生的成绩从键盘输入】
	//2. j 表示第几个班级
	//3. 定义一个变量存放总成绩

	//分析实现思路3
	//1. 我们可以把代码做活
	//2. 定义两个变量，表示班级的个数和班级的人数

	//统计三个班及格人数，每个班有5名同学
	//分析思路
	//1. 声明以变量 passCount 用于保存及格人数

	var classNum int
	var stuNum int
	fmt.Println("请输入班级数量：")
	fmt.Scanln(&classNum)
	fmt.Println("请输入班级人数")
	fmt.Scanln(&stuNum)
	var passCount int = 0
	var talSum float64
	for j:=1;j<=classNum;j++{
		var sum float64 = 0
		for i:=1;i<=stuNum;i++{
			var score float64 =0
			fmt.Printf("请输入%v班%v号学生的成绩",j,i)
			fmt.Scanln(&score)
			if score >60{
				passCount++
			}
			sum +=score
		}
		fmt.Printf("第班的平均成绩为%v \n",sum/float64(stuNum))
		talSum+=sum
	}
	fmt.Printf("所有班级的总平均分为：%v \n" ,talSum/float64(classNum*stuNum))
	fmt.Printf("及格人数为%v\n", passCount)
}