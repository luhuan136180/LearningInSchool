package model
//定义一个public构造体
type Student struct {
	Name string
	Age int
}
//定义一个private构造体
//因为student结构体首字母是小写，因此是只能在model使用
//我们通过工厂模式来解决
type student01 struct {
	Name string
	Age int
}
type student02 struct {
	name string
	Age int
}
//当其他包需要引用该构造体时，不能直接引用，需要通过函数（工厂模式解决）
func NewStudent(name string,age int) *student01 {
	return &student01{
		Name: name,
		Age: age,
	}
}
func NewStudent02(name string,age int) student01 {
	return student01{
		Name: name,
		Age: age,
	}
}
func NewStudent03(name string,age int) *student02 {
	return &student02{
		name: name,
		Age: age,
	}
}
func(stu *student02)GetName()string{
	return (*stu).name
}