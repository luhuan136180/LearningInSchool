package main
import "fmt"
/*第一题
  学生案例：
  编写一个Student结构体，包含name、gender、age、id、score字段，分别为string、string、int、int、float64类型。
  结构体中声明一个say方法，返回string类型，方法返回信息中包含所有字段值。
  在main方法中，创建Student结构体实例(变量)，并访问say方法，并将调用结果打印输出。

*/
type Student struct {
	name string
	gender string
	age int
	id int
	score float64
}
func(stu Student)say()string{
	inforStr:=fmt.Sprintf("student的信息name=[%v],gender=[%v],age=[%v],id=[%v],score=[%v]",
		stu.name,stu.gender,stu.age,stu.id,stu.score)
	return inforStr
}

//第二题：
//编写一个 Dog 结构体，包含 name、age、weight 字段
//结构体中声明一个 say 方法，返回 string 类型，方法返回信息中包含所有字段值。
// 在 main 方法中，创建 Dog 结构体实例(变量)，并访问 say 方法，将调用结果打印输出。
type Dog struct {
	name string
	age int
	weight int
}
func (dog Dog)say()string{
	infoStr:=fmt.Sprintf("dog的信息：name=[%v],age=[%v],weight=[%v]",
		dog.name,dog.age,dog.weight)
	return infoStr
}

//第三题
//编程创建一个 Box 结构体，在其中声明三个字段表示一个立方体的长、宽和高，长宽高要从终
//端获取
//2) 声明一个方法获取立方体的体积。
//3) 创建一个 Box 结构体变量，打印给定尺寸的立方体的体积
type Box struct{
	height float64
	weight float64
	length float64
}

func(box Box) PrintV() float64 {
	v:=box.length*box.height*box.weight
	return v
}
//第四题
//一个景区根据游人的年龄收取不同价格的门票，比如年龄大于 18，收费 20 元，
//其它情况门票免费.
//请编写 Visitor 结构体，根据年龄段决定能够购买的门票价格并输出
type Visitor struct{
	Name string
	Age int
}
func(visitor Visitor)showPrice(){
	Price:=0
	if visitor.Age>80||visitor.Age<5{
		fmt.Println("bushihe又玩")
	}else if visitor.Age>18{
		Price=20
		fmt.Printf("游客的名字为 %v 年龄为 %v 收费%v元 \n",visitor.Name,visitor.Age,Price)
	}else {
		fmt.Printf("游客的名字为 %v 年龄为 %v 收费%v元 \n",visitor.Name,visitor.Age,Price)
	}

}
func main(){
	stu01:=Student{
		name: "tom",
		gender: "nan",
		age:23,
		id:12345,
		score: 78.9,
	}
	fmt.Println(stu01.say())

	//
	dog01:=Dog{"a",4,7}
	fmt.Println(dog01.say())
	//
	var n1 Box
	fmt.Println("请输入长度：")
	fmt.Scanln(&n1.length)
	fmt.Println("请输入宽：")
	fmt.Scanln(&n1.weight)
	fmt.Println("请输入高度：")
	fmt.Scanln(&n1.height)
	fmt.Println("zhege boxde 体积为",n1.PrintV())


	var visitor01 Visitor
	for{
		fmt.Println("qingshuru你的名字：")
		fmt.Scanln(&visitor01.Name)
		if visitor01.Name=="n"{
			break
		}else{
			fmt.Println("请输入年龄:")
			fmt.Scanln(&visitor01.Age)
			visitor01.showPrice()
		}
	}
}
