package main

import (
	"fmt"
	"reflect" //实现运行时的反射
)

func reflectTest01(b interface{}) {
	//通过反射获取的传入的变量的 type , kind, 值
	//1. 先获取到 reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Println("RType=", rType)
	
	//2. 获取到 reflect.Value
	rVal := reflect.ValueOf(b)

	n2 := 2 + rVal.Int()
	fmt.Println("n2=", n2)

	//输出： 100 ， reflect.Value
	fmt.Printf("RVal=%v ;rVal Type=%T \n", rVal, rVal)

	//下面我们将 rVal 转成 interface{}
	iv := rVal.Interface()
	num2 := iv.(int)
	fmt.Println("num2=", num2)
}
func reflectTest02(b interface{}) {
	//通过反射获取的传入的变量的 type , kind, 值
	rType := reflect.TypeOf(b)
	fmt.Println("rType=", rType)

	rVal := reflect.ValueOf(b)

	//3. 获取 变量对应的Kind
	//(1) rVal.Kind() ==>
	kind1 := rVal.Kind()
	kind2 := rType.Kind()
	fmt.Printf("kind = %v;kind=%v \n", kind1, kind2)

	//下面我们将 rVal 转成 interface{}
	iV := rVal.Interface()
	fmt.Printf("iv=%v iv type=%T \n", iV, iV)
	//将 interface{} 通过断言转成需要的类型
	//这里，我们就简单使用了一带检测的类型断言.
	//同学们可以使用 swtich 的断言形式来做的更加的灵活
	stu, ok := iV.(Student)
	if ok {
		fmt.Printf("stu.Name=%v\n", stu.Name)
	}
}

type Student struct {
	Name string
	Age  int
}

type Monster struct {
	Name string
	Age  int
}

func main() {
	//请编写一个案例，
	//演示对(基本数据类型、interface{}、reflect.Value)进行反射的基本操作

	//1. 先定义一个int
	var num int = 100
	reflectTest01(num)

	//
	stu := Student{
		Name: "tom",
		Age:  20,
	}
	reflectTest02(stu)
}
