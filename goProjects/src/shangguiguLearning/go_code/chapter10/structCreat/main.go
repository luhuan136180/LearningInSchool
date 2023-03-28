package main
import "fmt"

type Person struct {
	Name string
	Age int
}

func main() {
	//方法一
	//案例演示: var person Person
	//前面我们已经说了。

	//方法二
	//案例演示: var person Person = Person{}
	p2 := Person{"mary",20}
	pp2 :=Person{}
	pp2.Age=20
	pp2.Name="mary"
	fmt.Println(p2)
	fmt.Println("pp2=",pp2)

	//方法3
	//案例: var person *Person = new (Person)
	var p3 *Person =new(Person)
	var pp3 *Person=new(Person)
	//因为p3是一个指针，因此标准的给字段赋值方式
	//(*p3).Name = "smith" 也可以这样写 p3.Name = "smith"

	//原因: go的设计者 为了程序员使用方便，底层会对 p3.Name = "smith" 进行处理
	//会给 p3 加上 取值运算 (*p3).Name = "smith"
	(*p3).Name="smith"
	(*p3).Age = 30

	pp3.Name = "smith"
	pp3.Age= 30
	fmt.Println("p3:",*p3)//输出：p3: {smith 30}
	fmt.Println("pp3:",pp3)//输出：pp3: &{smith 30}

	//方式4-{}
	//案例: var person *Person = &Person{}

	//下面的语句，也可以直接给字符赋值
	//var person *Person = &Person{"mary", 60}
	var person *Person = &Person{}
	var pperson *Person = &Person{"scott",88}
	//因为person 是一个指针，因此标准的访问字段的方法
	// (*person).Name = "scott"
	// go的设计者为了程序员使用方便，也可以 person.Name = "scott"
	// 原因和上面一样，底层会对 person.Name = "scott" 进行处理， 会加上 (*person)
	(*person).Name = "scott"
	person.Age=88
	fmt.Println("person:",*person)
	fmt.Println("pperson:",*pperson)



}
